package social

import (
	"testing"
)

// Test to unmarshal tweets from known file
// Test to get data from API

func TestSocialFetching(t *testing.T) {
	socialGrabber := NewSocialGrabber()

	go func() {
		for statuses := range socialGrabber.C {
			if len(statuses) == 0 {
				t.Error("Tweets are empty")
			}

			return
		}
	}()

	RunSocial(socialGrabber)

}
