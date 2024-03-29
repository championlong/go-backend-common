package utils

import (
	"context"
	"fmt"
	"runtime"
)

func Recovery() {
	if r := recover(); r != nil {
		buf := make([]byte, 1<<18)
		n := runtime.Stack(buf, false)
		fmt.Printf(" %+v. stack: %s", r, buf[0:n])
	}
}

func SafeGoroutine(f func()) {
	defer Recovery()
	f()
}

func SafeGoroutineByContext(ctx context.Context, f func(ctx context.Context)) {
	defer Recovery()
	f(ctx)
}
