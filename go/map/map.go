package main

import (
	"fmt"
)

type Status struct {
	St   string
	Data int
}

func main() {
	m := make(map[string]string)
	fmt.Println(len(m))
	if v, ok := m["12"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("no data")
	}
	tc := make(map[string]Status)
	v := tc["12"]
	fmt.Printf("%#v\n", v)

	m = map[string]string{"a": "1", "b": "2", "c": "3"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
