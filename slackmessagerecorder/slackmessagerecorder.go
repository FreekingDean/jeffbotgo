package slackmessagerecorder

import (
	"context"
	"log"

	"github.com/nlopes/slack/slackevents"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data *slackevents.MessageEvent
}

// HelloPubSub consumes a Pub/Sub message.
func Record(ctx context.Context, m PubSubMessage) error {
	log.Printf("%+v", m.Data)
	return nil
}
