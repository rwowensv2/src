package main

// still learning, error handling ugh
// https://gobyexample.com/command-line-flags
import (
	"errors"
	"flag"
	"fmt"
	"net"
	"time"
)

func hostcheck(h string, p string, s int64, x int64) string {
	timeout := time.Duration(x) * time.Second
	conn, err := net.DialTimeout("tcp", h+":"+p, timeout)

	if err != nil {
		err := errors.New("Timeout")
		// https://www.systutorials.com/241626/in-golang-how-to-convert-an-error-to-a-string/
		return err.Error()
	}
	return (conn.RemoteAddr().String())
}

func main() {

	hostPtr := flag.String("host", "golang.org", "hostname or IP")
	portPtr := flag.String("port", "80", "portnumber")
	sleepPtr := flag.Int64("sleep", 5, "Polling interval in seconds")
	timePtr := flag.Int64("timeout", 10, "TCP Timeout, default 10")
	flag.Parse()

	host := *hostPtr
	port := *portPtr
	sleep := *sleepPtr
	timeo := *timePtr
	fmt.Println("Connected:", hostcheck(host, port, sleep, timeo))

}
