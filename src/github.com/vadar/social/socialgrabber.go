package social

import (
	"fmt"
)

type SocialGrabber struct {
	C chan []Post
}

type Post struct {
	Message string
}

func NewSocialGrabber() *SocialGrabber {
	c := make(chan []Post)
	sg := &SocialGrabber{C: c}

	return sg
}

// TODO make channels/go routines for the workers
func RunSocial(callback *SocialGrabber) {
	fmt.Println("Getting social")

	twitterChan := GetTweets()

	select {
	case tweets := <-twitterChan:
		fmt.Printf("Got tweets %d\n", len(tweets))
		callback.C <- tweets
	}

	// callback.C <- statuses
}
