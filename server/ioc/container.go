package ioc

import (
	"fmt"
)

type Provider[T any] func() T

func Provide[T any](provider Provider[T]) {
	name := generateServiceName[T]()
	injector.set(name, provider())
}

func Invoke[T any]() T {
	name := generateServiceName[T]()
	obj := injector.get(name).(T)
	return obj
}

func generateServiceName[T any]() string {
	var t T
	// struct
	name := fmt.Sprintf("%T", t)
	if name != "<nil>" {
		return name
	}
	// interface
	return fmt.Sprintf("%T", new(T))
}
