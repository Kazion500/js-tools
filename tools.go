package toolkit

import (
	"encoding/json"
)

type AnyType interface {
}

type Tool[T any] interface {
	Stringify(any) (string, error)
	Parse(string, T) (T, error)
}

type tool[T AnyType] struct{}

func New[T AnyType]() Tool[T] {
	return &tool[T]{}
}

func (t *tool[T]) Stringify(data any) (string, error) {
	result, err := json.Marshal(data)
	return string(result), err
}

func (t *tool[T]) Parse(js string, v T) (T, error) {
	stringBytes := []byte(js)
	err := json.Unmarshal(stringBytes, &v)
	return v, err
}
