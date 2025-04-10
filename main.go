package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ordrid/orders-api/application"
)

func main() {
	app := application.New()

	if err := app.Start(context.TODO()); err != nil {
		fmt.Fprintln(os.Stderr, "failed to start app:", err)
	}
}
