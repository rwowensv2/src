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
        gop   = "GOPATH"
)

envvars := [3]string{user, shell, gop}

  for _, v := range envvars {
    fmt.Println(os.Getenv(v))
  }
}
