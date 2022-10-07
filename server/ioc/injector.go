package ioc

import "fmt"

type Injector struct {
	services map[string]any
}

func (i *Injector) exists(name string) bool {
	_, ok := i.services[name]
	return ok
}

func (i *Injector) set(name string, obj any) {
	if ok := i.exists(name); ok {
		panic(fmt.Errorf("DI: service `%s` has already been declared", name))
	}
	i.services[name] = obj
}

func (i *Injector) get(name string) any {
	if ok := i.exists(name); !ok {
		panic(fmt.Errorf("DI: service `%s` has not been declared", name))
	}
	return i.services[name]
}

var injector Injector = Injector{
	services: make(map[string]any),
}
