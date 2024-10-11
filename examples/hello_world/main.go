package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"

	"github.com/underopsco/go-action/pkg/action"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	fmt.Printf("Hello, %s!\n", action.GetInput("who-to-greet"))

	action.SetOutput("random-number", fmt.Sprint(rand.Int()))
}
