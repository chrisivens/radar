package textproc

import "github.com/vadar/social"

type MessageWorker struct {
	cake string
}

var stopwords []string

func init() {
	stopwords = []string{"and", "the", "it"} // FIXME - remove when data is fetching from the DB
}

func NewMessageProcessor() chan<- []social.Post {
	c := make(chan []social.Post)

	return c
}

// TODO - Get the stopwords from the DB
func StopWords() []string {
	return stopwords
}
