package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

func main() {

	//https://stackoverflow.com/questions/18973335/golang-check-number-of-arguments-also-user-input-check-for-return-key-blan
	redisServer := os.Args[1]
	password := os.Args[2]
	mastername := os.Args[3]

	//connr, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("cluster"))
	connr, err := redis.Dial("tcp", redisServer+":6379", redis.DialPassword(password))
	if err != nil {
		log.Fatal(err)
	}
	defer connr.Close()

	//conns, err := redis.Dial("tcp", "127.0.0.1:10002", redis.DialPassword("cluster"))
	conns, err := redis.Dial("tcp", redisServer+":10001", redis.DialPassword("changeme"))
	if err != nil {
		log.Fatal(err)
	}

	// Importantly, use defer to ensure the connection is always
	// properly closed before exiting the main() function.
	defer conns.Close()

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
	//role, err := redis.Strings(conns.Do("sentinel", "failover", mastername))
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
			conns.Send("sentinel", "failover", mastername)
			conns.Flush()
			a, b := conns.Receive()
			_ = b
			fmt.Println(a)
		}
	}
}
