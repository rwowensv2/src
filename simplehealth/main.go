package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {

	hostPtr := flag.String("host", "golang.org", "hostname or IP")
	portPtr := flag.String("port", "80", "portnumber")
	sleepPtr := flag.Int64("sleep", 5, "Polling interval in seconds")

	flag.Parse()

	for {
		conn, err := net.Dial("tcp", *hostPtr+":"+*portPtr)
		if err != nil {
			fmt.Println(err)
			return

		}

		fmt.Printf("Connection established between %s and localhost.\n", *hostPtr)
		fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
		fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())
		time.Sleep(time.Duration(*sleepPtr) * time.Second)

	}

}
