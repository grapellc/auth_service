package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-micro.dev/v5"
	"go-micro.dev/v5/health"

	"github.com/your-moon/grape-auth-service/internal/auth"
	"github.com/your-moon/grape-auth-service/internal/auth/transport"
	"github.com/your-moon/grape-auth-service/internal/otp"
	"github.com/your-moon/grape-shared/common/database"
	"github.com/your-moon/grape-shared/common/messaging"
)

const (
	defaultPort       = 8060
	defaultHealthPort = 8061
	serviceName       = "auth"
)

func main() {
	if err := loadConfig(); err != nil {
		log.Fatalf("config: %v", err)
	}

	initLogging()
	os.Setenv("MIGRATIONS_DIR", "./migrations")
	database.InitDB()
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("migrations: %v", err)
	}
	database.InitRedis()

	otpSender := initOTPSender()
	authSvc := auth.NewAuthService(database.DBClient, database.RedisClient, otpSender)
	handler := transport.NewAuthHandler(authSvc)

	port := viper.GetInt("auth_service.port")
	if port == 0 {
		port = defaultPort
	}
	healthPort := viper.GetInt("auth_service.health_port")
	if healthPort == 0 {
		healthPort = defaultHealthPort
	}

	addr := ":" + strconv.Itoa(port)
	svc := micro.NewService(
		micro.Name(serviceName),
		micro.Address(addr),
	)
	svc.Init()
	svc.Handle(handler)

	health.Register("database", health.PingCheck(func() error {
		sqlDB, err := database.DBClient.DB()
		if err != nil {
			return err
		}
		return sqlDB.Ping()
	}))
	health.Register("redis", health.CustomCheck(func() error {
		return database.RedisClient.Ping(context.Background()).Err()
	}))
	health.SetInfo("service", serviceName)

	healthMux := http.NewServeMux()
	health.RegisterHandlers(healthMux)
	healthServer := &http.Server{Addr: ":" + strconv.Itoa(healthPort), Handler: healthMux}
	go func() {
		if err := healthServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("health server: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-ctx.Done()
		stop()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = healthServer.Shutdown(shutdownCtx)
		os.Exit(0)
	}()

	logrus.Infof("auth-service listening on %s, health on :%d", addr, healthPort)
	if err := svc.Run(); err != nil {
		log.Fatalf("auth-service: %v", err)
	}
}

func loadConfig() error {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../") // Support local development from service root

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Explicitly bind environment variables for nested keys.
	// This is necessary because viper.AutomaticEnv() doesn't always pick up
	// environment variables for keys that aren't already in a config file.
	envBindings := []string{
		"db.host", "db.port", "db.user", "db.password", "db.dbname",
		"redis.host", "redis.port", "redis.password",
		"emqx.broker", "emqx.topic", "emqx.client_id",
		"smtp.host", "smtp.port", "smtp.username", "smtp.password", "smtp.from_email",
		"jwt.secret",
	}
	for _, key := range envBindings {
		if err := viper.BindEnv(key); err != nil {
			logrus.Warnf("Failed to bind env var for %s: %v", key, err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("Config file not found, using environment variables and defaults")
			return nil
		}
		return err
	}
	return nil
}

func initLogging() {
	lvl := viper.GetString("log.level")
	if lvl == "" {
		lvl = "info"
	}
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}

func initOTPSender() otp.Sender {
	var smsSender, emailSender otp.Sender
	if err := messaging.NewMqttClient(); err != nil {
		logrus.Errorf("OTP Sender: MQTT initialization failed: %v. SMS OTP will NOT be sent (falling back to console)", err)
		smsSender = otp.NewConsoleSender()
	} else {
		smsSender = otp.NewMqttSender(messaging.MqttClient, viper.GetString("emqx.topic"))
	}
	if viper.GetString("smtp.host") != "" {
		emailSender = otp.NewSMTPSender(
			viper.GetString("smtp.host"),
			viper.GetInt("smtp.port"),
			viper.GetString("smtp.username"),
			viper.GetString("smtp.password"),
			viper.GetString("smtp.from_email"),
		)
	} else {
		logrus.Errorf("OTP Sender: SMTP host not configured. Email OTP will NOT be sent (falling back to console)")
		emailSender = otp.NewConsoleSender()
	}
	return otp.NewCompositeSender(smsSender, emailSender)
}
