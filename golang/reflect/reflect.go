package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	fmt.Println("Type: ", reflect.TypeOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type())
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("value: ", v.Float())

	p := reflect.ValueOf(&x)
	fmt.Println("type of p: ", p.Type())
	fmt.Println("settability of p :", p.CanSet())

	v = p.Elem()
	fmt.Println("settability of v: ", v.CanSet())

	v.SetFloat(7.2)
	fmt.Println(v.Interface())
	fmt.Println(x)

	t := T{345, "smile123"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
