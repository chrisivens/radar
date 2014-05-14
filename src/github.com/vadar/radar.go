package main

import (
	"fmt"
	"github.com/vadar/social"
	"github.com/vadar/textproc"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type StopWords struct {
	Id    bson.ObjectId `json:"id"    bson:"_id,omitempty"`
	Words []string      `json:"words" bson:"words"`
}

// type (
// 	radarRepo struct {
// 		Collection *mgo.Collection
// 	}
// )

func main() {
	ticker := time.Tick(time.Second * 5)
	social_grabber := social.NewSocialGrabber()
	var message_processor chan<- []social.Post = textproc.NewMessageProcessor()

	// var (
	// 	mongoSession *mgo.Session
	// 	database     *mgo.Database
	// 	collection   *mgo.Collection
	// 	err          error
	// )

	// if mongoSession, err = mgo.Dial("localhost"); err != nil {
	// 	panic(err)
	// }
	// database = mongoSession.DB("radar")
	// collection = database.C("stopwords")

	// fmt.Printf("Collection: %# v", collection)

	for {
		select {
		case statuses := <-social_grabber.C:
			fmt.Printf("Got statuses: %d\n", len(statuses))
			message_processor <- statuses

		case <-ticker:
			go social.RunSocial(social_grabber)
		}
	}

}
