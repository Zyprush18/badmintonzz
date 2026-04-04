package cntx

import (
	"context"
	"time"
)

var (
	DbShortTime time.Duration = 5 * time.Second
	DBLongTime time.Duration = 10 * time.Second
)

func TimeoutShortContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, DbShortTime)
}


func TimeoutLongContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, DBLongTime)
}