package action

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	inputField  fieldKind = "input"
	outputField           = "output"
)

type fieldKind string

type tagParts struct {
	name string
	kind fieldKind
}

func bindInputs(i interface{}) error {
	return visitFields(i, func(f reflect.Value, tag tagParts) error {
		switch tag.kind {
		case inputField:
			f.SetString(GetInput(tag.name))
		}

		return nil
	})
}

func bindOutputs(i interface{}) error {
	return visitFields(i, func(f reflect.Value, tag tagParts) error {
		switch tag.kind {
		case outputField:
			SetOutput(tag.name, fmt.Sprintf("%d", f.Int()))
		}

		return nil
	})
}

func visitFields(i interface{}, visitor func(f reflect.Value, tag tagParts) error) error {
	t := reflect.TypeOf(i).Elem()
	v := reflect.ValueOf(i).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		tag := t.Field(i).Tag.Get("action")

		if !f.IsValid() || !f.CanSet() || tag == "" {
			continue
		}

		if tp := parseTag(tag); tp != nil {
			if err := visitor(f, *tp); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseTag(tag string) *tagParts {
	parts := strings.Split(tag, ",")

	if len(parts) == 0 {
		return nil
	}

	tagParts := &tagParts{
		name: parts[0],
		kind: inputField,
	}

	if len(parts) > 1 {
		tagParts.kind = fieldKind(parts[1])
	}

	return tagParts
}
