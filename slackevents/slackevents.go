package slackevents

import (
	"bytes"
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
)

func SlackEvent(w http.ResponseWriter, r *http.Request) {
	_, err := pubsub.NewClient(context.Backgound(), "project-id")
	if err != nil {
		// TODO: Handle error.
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	body := buf.String()
	log.Println(body)
	w.WriteHeader(200)
}
