package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
	fir  int

	Name string `abc:"number", def:"231"`
}

func (this *MyStruct) GetName() string {
	return this.name
}

func main() {
	s := "This is string"
	var x float64 = 3.4
	a := new(MyStruct)
	a.name = "shilezhi"
	ss := []string{"1", "2", "3"}

	fmt.Printf("ValueOf: %v\n", reflect.ValueOf(s))
	fmt.Printf("       : %v\n", reflect.ValueOf(x))
	fmt.Printf("       : %v\n", reflect.ValueOf(a))
	fmt.Printf("       : %v\n", reflect.ValueOf(ss))
	fmt.Println("-------------------")

	fmt.Printf("TypeOf: %v\n", reflect.TypeOf(s))
	fmt.Printf("      : %v\n", reflect.TypeOf(x))
	fmt.Printf("      : %v\n", reflect.TypeOf(a))
	fmt.Printf("      : %v\n", reflect.TypeOf(ss))
	fmt.Println("-------------------")

	fmt.Printf("NumMethod: %v\n", reflect.TypeOf(s).NumMethod())
	fmt.Printf("         : %v\n", reflect.TypeOf(x).NumMethod())
	fmt.Printf("         : %v\n", reflect.TypeOf(a).NumMethod())
	fmt.Printf("         : %v\n", reflect.TypeOf(ss).NumMethod())
	fmt.Println("-------------------")

	fmt.Printf("Kind: %v\n", reflect.ValueOf(s).Kind())
	fmt.Printf("    : %v\n", reflect.ValueOf(x).Kind())
	fmt.Printf("    : %v\n", reflect.ValueOf(a).Kind())
	fmt.Printf("    : %v\n", reflect.ValueOf(ss).Kind())
	fmt.Println("-------------------")

	am := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Printf("MethodByName: %v\n", am[0])
	fmt.Println("-------------------")

	aty := reflect.TypeOf(*a)
	for i := 0; i < aty.NumField(); i++ {
		fmt.Printf("Tag: '%s'\n", aty.Field(i).Tag.Get("abc"))
	}
}
