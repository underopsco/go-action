package main

import (
	"fmt"
	"math/rand"

	"github.com/wtfiscrq/go-action/pkg/action"
)

type HelloWorld struct {
	WhoToGreet   string `action:"who-to-greet"`
	RandomNumber int    `action:"random-number,output"`
}

func (hw *HelloWorld) Run() error {
	fmt.Printf("Hello, %s!\n", hw.WhoToGreet)

	hw.RandomNumber = rand.Int()

	return nil
}

func main() {
	if err := action.Execute(&HelloWorld{}); err != nil {
		action.SetFailed(err, map[string]string{})
	}
}
