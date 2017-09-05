package main

import (
	"fmt"
	"reflect"
)

func map2struct(a map[string]interface{}, result interface{}) error {
	vv := reflect.ValueOf(result)
	if vv.Kind() != reflect.Ptr {
		return fmt.Errorf("result must a ptr")
	}
	vv = vv.Elem()
	vt := reflect.TypeOf(result).Elem()
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		key := field.Tag.Get("xx")
		if v, ok := a[key]; ok {
			vv.FieldByName(field.Name).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func main() {
	type MyStruct struct {
		Name string `xx:"name"`
		Age  int    `xx:"age"`
	}
	a := make(map[string]interface{})
	a["name"] = "Tony"
	a["age"] = 23

	result := new(MyStruct)
	// var result MyStruct
	if err := map2struct(a, result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
