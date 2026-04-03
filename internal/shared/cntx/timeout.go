package cntx

import (
	"context"
	"time"
)

var (
	DbShortTime int = 3
	DbMediumTime int = 5
	DBLongTime int = 10
)

func TimeOutContext(ctx context.Context, duration int) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, time.Duration(duration) * time.Second)
}