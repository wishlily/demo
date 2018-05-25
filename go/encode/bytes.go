package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//  结构体需要大写命名（对外可见，否则报错）
	type MyStruct struct {
		A    uint32
		Int  int32
		B    float32
		Bool bool
	}
	a := MyStruct{
		9999,
		2333,
		66,
		true,
	}
	fmt.Printf("DATA:   %v\n", a)
	buf := &bytes.Buffer{}
	// struct 中包含
	// int, string, []byte 无法使用：invalid type
	if err := binary.Write(buf, binary.BigEndian, a); err != nil {
		panic(err)
	}
	fmt.Printf("BYTES:  %v\n", buf.Bytes())
	v := &MyStruct{}
	if err := binary.Read(buf, binary.BigEndian, v); err != nil {
		panic(err)
	}
	fmt.Printf("RESULT: %v\n", v)
}
