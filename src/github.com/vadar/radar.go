package main

import (
	"fmt"
	"github.com/vadar/social"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")
	social.SocialGrabber()

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
