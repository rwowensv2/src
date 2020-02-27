package main

import (
  "fmt"
  "os"
//  "strings"
)

func main() {

const (
	robVAR = "ROBVAR"
)

fmt.Println("My Env Var:", os.Getenv(robVAR))
	
}
