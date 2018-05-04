package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type GobA struct {
	Str string
	Int int
}

type GobB struct {
	Str  string
	Int  int
	Flot float32
	Mty  GobA
	Bool bool
}

func main() {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	dec := gob.NewDecoder(&buff)

	{
		if err := enc.Encode(GobA{"GobA", 1}); err != nil {
			panic(err)
		}
		fmt.Printf("%d,%v\n", buff.Len(), buff.Bytes())

		var v GobB
		if err := dec.Decode(&v); err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", v)
	}

	{
		if err := enc.Encode(GobB{"GobB", 2, 3.33, GobA{"G", 3}, true}); err != nil {
			panic(err)
		}
		fmt.Printf("%d,%v\n", buff.Len(), buff.Bytes())

		var v GobA
		if err := dec.Decode(&v); err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", v)
	}
}
