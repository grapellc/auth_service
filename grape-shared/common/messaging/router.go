package messaging

import (
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

var (
	Router *message.Router
)

func NewRouter() (*message.Router, error) {
	if Router != nil {
		return Router, nil
	}

	router, err := message.NewRouter(message.RouterConfig{}, Logger)
	if err != nil {
		return nil, err
	}

	router.AddMiddleware((&middleware.DelayOnError{
		InitialInterval: time.Second * 1,
		MaxInterval:     time.Second * 10,
		Multiplier:      2.0,
	}).Middleware)

	Router = router
	return router, nil
}
