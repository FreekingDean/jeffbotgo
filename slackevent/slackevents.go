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
	_, err := pubsub.NewClient(context.Background(), "jeffbot")
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
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body))
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			log.Println(ev.Text)
			//api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		}
	}
	log.Println(body)
	w.WriteHeader(200)
}
