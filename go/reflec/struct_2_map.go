package main

import (
	"fmt"
	"reflect"
)

func struct2map(v interface{}) (map[string]interface{}, error) {
	vv := reflect.ValueOf(v)
	vt := reflect.TypeOf(v)
	if vv.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}
	m := make(map[string]interface{})
	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		k := f.Tag.Get("xx")
		if len(k) != 0 {
			m[k] = vv.FieldByName(f.Name).String()
		}
	}
	return m, nil
}

func main() {
	type VehicleInfo struct {
		vehicleId string `xx:"number"`
		Date      string `xx:"date"`
		Type      string `xx:"type"`
		Brand     string `xx:"brand"`
		Color     string `xx:"color"`
		haha      []byte
	}
	a := &VehicleInfo{
		vehicleId: "123456",
		Date:      "20140101",
		Type:      "Truck",
		Brand:     "Ford",
		Color:     "White",
		haha:      []byte{1, 2, 3},
	}
	if v, err := struct2map(a); err != nil {
		panic(err)
	} else {
		fmt.Println(v)
	}
}
