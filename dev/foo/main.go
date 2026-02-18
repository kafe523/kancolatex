package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	E int
}

func (f *Foo) A() int {
	return 69
}

func (f Foo) B() int {
	return 420
}

func main() {
	var foo Foo

	rv := reflect.ValueOf(&foo)
	rt := reflect.TypeOf(&foo)
	fmt.Printf("%#v\n", rv)
	fmt.Printf("%#v\n", rt)

	fmt.Printf("%#v\n", rv.NumMethod())
	fmt.Printf("%#v\n", rt.NumMethod())

	for i := range rt.NumMethod() {
		m := rt.Method(i)
		fmt.Printf("%#v\n", m)
		result := m.Func.Call([]reflect.Value{rv})
		if len(result) != 1 {
			panic("only support 1 argument method return")
		}
		fmt.Printf("%#v\n", result[0].Interface())
	}
}
