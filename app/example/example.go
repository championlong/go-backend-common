package main

import (
	"fmt"
	"go-backend-common/app"
)

func main() {
	appInit := app.NewApp("Mock Server",
		"mock-basename",
		app.WithDescription("mock描述信息"),
		app.WithRunFunc(run()),
		app.WithNoConfig(),
		app.WithCommand(app.NewCommand("ccc", "ddd")),
	)
	appInit.Run()
}

func run() app.RunFunc {
	return func(basename string) error {
		fmt.Println("init mock app")
		for {
		}
		return nil
	}
}
