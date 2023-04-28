package jskit

import (
	"encoding/json"
)

type AnyType interface{}

type JSKit[T AnyType] interface {
	JSON() JSON[T]
	ToString(value []byte) string
}

type JSON[T AnyType] interface {
	Stringify(any) (string, error)
	Parse(string, T) (T, error)
}

type jsKit[T AnyType] struct{}

type jsont[T AnyType] struct{}

func NewJSKit[T AnyType]() JSKit[T] {
	return &jsKit[T]{}
}

func (js *jsKit[T]) JSON() JSON[T] {
	return &jsont[T]{}
}

func (j *jsont[T]) Stringify(data any) (string, error) {
	jskit := NewJSKit[T]()
	result, err := json.Marshal(data)
	return jskit.ToString(result), err
}

func (j *jsont[T]) Parse(js string, v T) (T, error) {
	stringBytes := []byte(js)
	err := json.Unmarshal(stringBytes, &v)
	return v, err
}

func (js *jsKit[T]) ToString(value []byte) string {
	return string(value)
}
