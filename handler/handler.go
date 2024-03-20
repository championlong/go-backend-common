package handler

import (
	"go-backend-common/handler/service"
	"sync"
)

var (
	initOnce sync.Once
)

func init() {
	initOnce.Do(func() {
		service.NewService()
	})
}
