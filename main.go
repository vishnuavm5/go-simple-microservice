package main

import (
	"context"
	"fmt"

	"github.com/vishnuavm5/orders-api/application"
)

func main() {
	app := application.New()
	fmt.Println("Server is running")
	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
