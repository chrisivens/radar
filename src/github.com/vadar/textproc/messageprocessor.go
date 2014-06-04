package textproc

import (
	"fmt"
	"github.com/vadar/nlp"
	"github.com/vadar/social"
)

type MessageWorker struct {
	cake string
}

type Message struct {
	*social.Post

	TF map[string]float32
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
			fmt.Println(tokens)
			message := new(Message)
			message.Post = post
			message.TF = nlp.TermFrequency(tokens)

			fmt.Println(message)
		}

	}
}
