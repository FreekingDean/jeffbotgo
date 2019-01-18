package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	//"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

func SlackEvent(w http.ResponseWriter, r *http.Request) {
	//slackToken := os.Getenv("SLACK_TOKEN")
	//api := slack.New(slackToken)
	pubsubClient, err := pubsub.NewClient(context.Background(), "jeffbot")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	topic := pubsubClient.Topic("slack")
	defer topic.Stop()
	defer pubsubClient.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
	}
	body := buf.String()
	eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body))
	//, slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: slackToken}))
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
		w.WriteHeader(200)
		return
	}
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			_ = topic.Publish(context.Background(), &pubsub.Message{Data: []byte(ev.Text)})
		case *slackevents.AppMentionEvent:
			//api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
		}
	}
	w.WriteHeader(200)
}
