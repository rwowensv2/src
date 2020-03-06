package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("cluster"))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	role, err := redis.Strings(conn.Do("ROLE"))
	if err != nil {
		// handle error
	}
	for _, key := range role {
		//	fmt.Println(key)
		if key == "master" {
			fmt.Println("it's the master")
			return
		} else if key == "slave" {
			fmt.Println("its' the slave")
			return
		}
	}
}
