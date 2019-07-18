package main

// https://gobyexample.com/command-line-flags
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
	timePtr := flag.Int64("timeout", 10, "TCP Timeout, default 10")
	flag.Parse()

	for {
		timeout := time.Duration(*timePtr) * time.Second
		conn, err := net.DialTimeout("tcp", *hostPtr+":"+*portPtr, timeout)

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
