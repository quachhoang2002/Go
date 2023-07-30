package log

import (
	"context"
)

// Logger methods interface
type Logger interface {
	Debug(ctx context.Context, args ...any)
	Debugf(ctx context.Context, template string, args ...any)
	Info(ctx context.Context, args ...any)
	Infof(ctx context.Context, template string, args ...any)
	Warn(ctx context.Context, args ...any)
	Warnf(ctx context.Context, template string, args ...any)
	Error(ctx context.Context, args ...any)
	Errorf(ctx context.Context, template string, args ...any)
	DPanic(ctx context.Context, args ...any)
	DPanicf(ctx context.Context, template string, args ...any)
	Fatal(ctx context.Context, args ...any)
	Fatalf(ctx context.Context, template string, args ...any)
}
