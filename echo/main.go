package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for min, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Printf("%x %s\n", min, arg)
	}
	//fmt.Println(s)

}
