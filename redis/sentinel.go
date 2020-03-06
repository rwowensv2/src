package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

func main() {

	connr, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("cluster"))
	if err != nil {
		log.Fatal(err)
	}
	defer connr.Close()

	conns, err := redis.Dial("tcp", "127.0.0.1:10002", redis.DialPassword("cluster"))
	if err != nil {
		log.Fatal(err)
	}

	defer conns.Close()

	//        conn.Send("ROLE")
	//       if err != nil {
	//             	log.Fatal(err)
	//	}
	//        conn.Send("ROLE")
	//        conn.Send("GET", "foo")
	//        conn.Flush()
	//        one, two := conn.Receive()
	//        fmt.Println("1: ",one,"2: ",two)
	//	s, err := redis.String(conn.Do("GET", "foo"))
	//	s, err := redis.String(conn.Do("KEYS", "*"))
	//	fmt.Printf("%#v\n", s)
	//
	//role, err := redis.Strings(conn.Do("KEYS", "*"))
	role, err := redis.Strings(connr.Do("ROLE"))
	//role, err := redis.Strings(conns.Do("sentinel", "failover", "OWENS"))
	if err != nil {
		// handle error
	}
	for _, key := range role {
		//	fmt.Println(key)
		if key == "master" {
			fmt.Println("Good Master")
			return
		} else if key == "slave" {
			fmt.Println("Slave Failover initiated...")
			conns.Send("sentinel", "failover", "OWENS")
			conns.Flush()
			a, b := conns.Receive()
			_ = b
			fmt.Println(a)
		}
	}
}
