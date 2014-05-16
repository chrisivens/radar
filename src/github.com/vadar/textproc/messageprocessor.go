package textproc

import (
	"github.com/vadar/social"
)

type MessageWorker struct {
	cake string
}

var stopwords []string

func init() {
	// FIXME - remove when data is fetching from the DB
	stopwords = []string{"and", "the", "it", "are", "is", "of", "at", "an"}
}

func NewMessageProcessor() chan<- []social.Post {
	c := make(chan []social.Post)

	go messageListen(c)

	return c
}

func messageListen(channel <-chan []social.Post) {
	for {
		var posts []social.Post = <-channel

		for i := 0; i < len(posts); i++ {
			var post *social.Post = &posts[i]
			tokens, fields := Tokenise(post.Message)

			post.Tokens = tokens
			post.Fields = fields
		}

	}
}
