package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	result := auditSearch()
	fmt.Printf("{\"loginfail\": %s}", result)
}

func auditSearch() string {
	curtime := time.Now()
	date := curtime.Format("01/02/2006")
	count := 30
	then := curtime.Add(time.Duration(-count) * time.Second)
	time := then.Format("15:04:05")

	var cmd, s, sep, e string
	s = "sudo ausearch --start "
	e = " -ua 0 --success no -i 2> /dev/null | grep isec | grep -v \"pubkey\\|gssapi\" | grep -v ^- | wc -l"
	sep = " "
	cmd += s + date + sep + time + e
	//fmt.Printf("%s\n", cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Failed to execute command: %s", cmd)
	}
	out = out[:len(out)-1]
	return string(out)
}
