package main

import (
	"fmt"
	"github.com/vadar/social"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")

	ticker := time.Tick(time.Second * 5)
	socialGrabber := social.NewSocialGrabber()

	for {
		select {
		case statuses := <-socialGrabber.C:
			fmt.Printf("Got statuses: %d\n", len(statuses))

		case <-ticker:
			go social.RunSocial(socialGrabber)
		}
	}

}
