package config

import (
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Init sets up viper (config file + env) and default log level.
// Services must call database.InitDB(), messaging, etc. themselves.
func Init() {
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config: %v", err)
		}
	}

	logrus.SetOutput(os.Stdout)
	logLevel := viper.GetString("log.level")
	if logLevel == "" {
		logLevel = "info"
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}
}
