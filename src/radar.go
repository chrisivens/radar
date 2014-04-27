package main

import (
	"fmt"
	"github.com/vadar"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")
	vadar.SocialGrabber()
	ticker := time.NewTicker(time.Second * 2)
	timer := time.NewTimer(time.Second * 10)

	for {
		select {
		case <-ticker.C:
			fmt.Println("woop")
		case <-timer.C:
			ticker.Stop()
			return
		}
	}

}
