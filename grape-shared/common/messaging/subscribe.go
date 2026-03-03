package messaging

import (
	"context"
)

func Subscribe() {
	Router.Run(context.Background())
}
