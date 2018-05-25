package main

import (
	"container/list"
	"fmt"
)

type One interface {
	Fun1(int) string
	Fun2(int) string
}

type Test1 struct {
	flag int
}

func (m *Test1) Fun1(a int) string {
	return fmt.Sprint("Test1 Fun1", a)
}

func (m *Test1) Fun2(a int) string {
	return fmt.Sprint("Test1 Fun2", a)
}

type Test2 struct {
	flag int
}

func (m *Test2) Fun1(a int) string {
	return fmt.Sprint("Test2 Fun1", a)
}

func (m *Test2) Fun2(a int) string {
	return fmt.Sprint("Test2 Fun2", a)
}

var l *list.List

func main() {
	t1 := new(Test1)
	t1.flag = 1
	t2 := new(Test2)
	t2.flag = 2
	t3 := new(Test1)
	t3.flag = 3
	t4 := new(Test2)
	t4.flag = 4
	l = list.New()
	Regist(t1)
	Regist(t2)
	Regist(t3)
	Regist(t4)

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		i++
		fmt.Println("======================")
		fmt.Printf("%#v\n", e.Value)
		// if v, ok := e.Value.(One); ok {
		// 	fmt.Println(v.Fun1(i))
		// 	fmt.Println(v.Fun2(i))
		// }
	}
	UnRegist(t1)
	fmt.Println()
	for e := l.Front(); e != nil; e = e.Next() {
		i++
		fmt.Println("======================")
		fmt.Printf("%#v\n", e.Value)
		// if v, ok := e.Value.(One); ok {
		// 	fmt.Println(v.Fun1(0))
		// 	fmt.Println(v.Fun2(0))
		// }
	}
}

func Regist(a One) error {
	l.PushBack(a)
	return nil
}

func UnRegist(a One) error {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == a {
			l.Remove(e)
		}
	}

	return nil
}
