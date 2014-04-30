package social

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Tweet struct {
	Coordinates  string        `json:"coordinates"`
	Favourited   bool          `json:"favorited"`
	Truncated    bool          `json:"truncated"`
	Created      string        `json:"created_at"`
	IdString     string        `json:"id_str"`
	Entities     TwitterEntity `json:"entities"`
	ReplyTo      string        `json:"in_reply_to_user_id_str"`
	Contributors string        `json:"contributors"`
	Text         string        `json:"text"`

	Metadata struct {
		Language   string `json:"iso_language_code"`
		ResultType string `json:"result_type"`
	} `json:"metadata"`

	RetweetCount      int                    `json:"retweet_count"`
	ReplyToStatus     string                 `json:"in_reply_to_status_id_str"`
	Id                int                    `json:"id"`
	Geo               string                 `json:"geo"`
	Retweeted         bool                   `json:"retweeted"`
	ReplyToUser       string                 `json:"in_reply_to_user_id"`
	Place             string                 `json:"place"`
	User              map[string]interface{} `json:"user"`
	ReplyToScreenName string                 `json:"in_reply_to_screen_name"`
	Source            string                 `json:"source"`
	ReplyToStatusId   string                 `json:"in_reply_to_status_id"`
}

type TwitterEntity struct {
	Urls         []string   `json:"urls"`
	Hashtags     []*Hashtag `json:"hashtags"`
	UserMentions []string   `json:"user_mentions"`
}

type Hashtag struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
}

type Statuses struct {
	Tweets []*Tweet `json:"statuses"`
}

func GetTweets() chan []Post {
	twitterChan := make(chan []Post)

	go TweetFetch(twitterChan)

	return twitterChan
}

func TweetFetch(out chan []Post) {
	path := os.Getenv("GOPATH")

	fi, err := os.Open(path + "/src/github.com/vadar/social/fixtures/tweets.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Opened file")

	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	statuses, err := UnmarshalTweets(fi)

	if err != nil {
		panic(err)
	}

	posts := make([]Post, len(statuses.Tweets))

	for i, tweet := range statuses.Tweets {
		posts[i] = Post{tweet.Text}
	}

	out <- posts
}

func UnmarshalTweets(reader io.Reader) (Statuses, error) {
	decoder := json.NewDecoder(reader)

	var tweets Statuses

	for {
		if err := decoder.Decode(&tweets); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return tweets, nil

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
