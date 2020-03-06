package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "10.3.234.100:10002", redis.DialPassword("cluster"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	role, err := redis.Strings(conn.Do("sentinel", "master", "OWENS"))
	if err != nil {
		// handle error
	}
	for _, key := range role {
		fmt.Println(key)
	}
}

