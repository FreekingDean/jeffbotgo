package slackevent

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/nlopes/slack/slackevents"
)

func SlackEvent(w http.ResponseWriter, r *http.Request) {
	client, err := pubsub.NewClient(context.Background(), "jeffbot")
	if err != nil {
		// TODO: Handle error.
	}
	topic := client.Topic("slack")

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	body := buf.String()
	log.Println(body)
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			log.Println(ev.Text)
			data, err := json.Marshal(ev)
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}
			_, err = topic.Publish(context.Background(), &pubsub.Message{Data: []byte(data)}).Get(context.Background())
			if err != nil {
				log.Println(err)
				w.WriteHeader(500)
				return
			}
		}
	}
	log.Println(body)
	w.WriteHeader(200)
}
