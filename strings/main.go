package main

import (
	"fmt"
)

func main() {

	//mystring := "hello, world"
	s := "hello, world"
	//s := []rune(mystring)

	fmt.Println(s[0:5])

	myshort := string(s[0:4])
	fmt.Println(myshort)

	fmt.Println("goodbye " + s[0:4])

	fmt.Println(s[2])
}
