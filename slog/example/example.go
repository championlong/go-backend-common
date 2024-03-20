package main

import (
	"context"
	"github.com/championlong/go-backend-common/slog"
)

func main() {
	slog.AddContextHook(ServiceContextHook)
	slog.WithContext(context.Background()).Debugf("1111")
	slog.Debugf("1111")
}

func ServiceContextHook(ctx context.Context) []slog.Field {
	var fields []slog.Field
	if ctx != nil {
		fields = append(fields, slog.String("test", "测试信息"))
	}
	return fields
}
