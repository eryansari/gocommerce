package repositories

import (
	"context"
	"time"

	"github.com/spf13/viper"
)

// Err
const (
	ErrMgoOpsFail = "failed get operation : %s"
)

// Set Ctx Timeout
func TimeOutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(viper.GetInt64("app.mongodb.timeoutinsecond"))*time.Second)
}

// Set Ctx Timeout Parent Assosicate
func TimeOutContextWithParent(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, time.Duration(viper.GetInt64("app.mongodb.timeoutinsecond"))*time.Second)
}
