package social

import "testing"
import "github.com/vadar/social"

// Test to unmarshal tweets from known file
// Test to get data from API

func TestSocialFetching(t *testing.T) {
	statuses, err := social.SocialGrabber()

	if err != nil {
		t.Error("Tweet status not fetching correctly")
	}

	if len(statuses.Tweets) == 0 {
		t.Error("Tweets are empty")
	}
}
