package repositories

import (
	"context"
	"github.com/spf13/viper"
	"time"
)

const (
	ErrMgoOpsFail = "failed get operation : %s"
)

// Set Ctx Timeout Parent Assosicate
func TimeOutContextWithParent(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, time.Duration(viper.GetInt64("app.mongodb.timeoutinsecond"))*time.Second)
}
