package main

import (
  "fmt"
  "os"
//  "strings"
)

func main() {

const (
	user  = "USER"
        shell = "SHELL"
)

envvars := [2]string{user, shell}

  for _, v := range envvars {
    fmt.Println(os.Getenv(v))
  }
}
