package main

import (
	"fmt"
	"math/rand"

	"github.com/underopsco/go-action/pkg/action"
)

func main() {
	fmt.Printf("Hello, %s!\n", action.GetInput("who-to-greet"))

	action.SetOutput("random-number", fmt.Sprint(rand.Int()))
}
