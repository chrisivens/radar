package main

import (
	"fmt"
	"socialgrabber"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")

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
