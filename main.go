package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	pinger("google.com")
}

func pinger(host string) {
	for {
		cmd := exec.Command("ping", "-c", "1", host)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Ping %s failed: %s\n", host, err)
		} else {
			fmt.Printf("Ping %s succeeded\n", host)
		}
		time.Sleep(1 * time.Second)
	}
}