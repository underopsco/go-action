package action

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

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
		return nil
	}

	inputs := map[string]string{}
	if err := json.Unmarshal([]byte(*raw), &inputs); err != nil {
		return err
	}

	for k, v := range inputs {
		envKey := fmt.Sprintf("INPUT_%s", strings.ToUpper(strings.Replace(k, " ", "_", -1)))

		if err := os.Setenv(envKey, v); err != nil {
			return err
		}
	}

	return nil
}
