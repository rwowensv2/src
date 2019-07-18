package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//args := os.Args[1:]
	//fmt.Println(len(os.Args))

	if len(os.Args) != 3 {
		fmt.Printf("usage: portcheck <hostname> <port>\nMissing port defaults to 80\n")
		return
	}

	//fmt.Println(args)
	hostname := os.Args[1]
	portnum := os.Args[2]
	//fmt.Println(hostname)
	//fmt.Println(portnum)

	conn, err := net.Dial("tcp", hostname+":"+portnum)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Connection established between %s and localhost.\n", hostname)
	fmt.Printf("Remote Address : %s \n", conn.RemoteAddr().String())
	fmt.Printf("Local Address : %s \n", conn.LocalAddr().String())
}
