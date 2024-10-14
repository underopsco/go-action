package action

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/google/go-github/v66/github"
)

// Deprecated: use github.NewClient(nil).WithAuthToken(Context.Token) instead
var REST = github.NewClient(nil).WithAuthToken(Context.Token)

type Action interface {
	Run() error
}

func Execute(a Action) error {
	// workaround for composite actions
	// see https://github.com/actions/runner/issues/665
	if err := loadInputsFromFlag(); err != nil {
		return err
	}

	if err := bindInputs(a); err != nil {
		return err
	}

	err := a.Run()

	if err := bindOutputs(a); err != nil {
		return err
	}

	return err
}

// workaround for composite actions
// see https://github.com/actions/runner/issues/665
func loadInputsFromFlag() error {
	raw := flag.String("inputs", "{}", "GitHub Action inputs")
	flag.Parse()
	if raw == nil {
		slog.Debug("no inputs provided")
		return nil
	}

	inputs := map[string]string{}
	if err := json.Unmarshal([]byte(*raw), &inputs); err != nil {
		return err
	}
	slog.Debug("inputs flag provided", slog.Any("value", inputs))

	for k, v := range inputs {
		envKey := fmt.Sprintf("INPUT_%s", strings.ToUpper(strings.Replace(k, " ", "_", -1)))

		if err := os.Setenv(envKey, v); err != nil {
			return err
		}
	}

	return nil
}
