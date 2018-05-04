package main

import (
	"flag"
	"fmt"
)

var svar string

// cmdlineargs -word=voo -numb=30 ...
func main() {
	flag.StringVar(&svar, "svar", "bar", "a string var")
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
