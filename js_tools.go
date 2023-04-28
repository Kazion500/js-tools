package js_tools

import (
	"encoding/json"
)

type AnyType interface {
}

type JSON[T any] interface {
	Stringify(any) (string, error)
	Parse(string, T) (T, error)
}

type jsont[T AnyType] struct{}

func NewJSON[T AnyType]() JSON[T] {
	return &jsont[T]{}
}

func (t *jsont[T]) Stringify(data any) (string, error) {
	result, err := json.Marshal(data)
	return string(result), err
}

func (t *jsont[T]) Parse(js string, v T) (T, error) {
	stringBytes := []byte(js)
	err := json.Unmarshal(stringBytes, &v)
	return v, err
}
