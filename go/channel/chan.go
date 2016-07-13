package main

import (
	"fmt"
)

type person struct {
	name string
	age  uint
}

type store struct {
	buffer chan person
}

func (m *store) get() string {
	ss := ""
	ok := true
	for ok {
		select {
		case v := <-m.buffer:
			ss += fmt.Sprintf("%v-%v,", v.name, v.age)
		default:
			ok = false
		}
	}
	return ss
}

func (m *store) put(v person) error {
	select {
	case m.buffer <- v:
		return nil
	default:
	}
	return fmt.Errorf("Buffer is full")
}

func main() {
	ptr := new(store)
	ptr.buffer = make(chan person, 3)
	persons := []person{
		{"Charles", 17},
		{"Erik", 17},
		{"McAvoy", 1},
		{"Fassbender", 3},
	}
	for _, v := range persons {
		if err := ptr.put(v); err != nil {
			fmt.Println(err, v)
		}
	}
	ss := ptr.get()
	fmt.Println(ss)
}
