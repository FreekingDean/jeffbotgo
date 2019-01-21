package slackmessagerecorder

import (
	"context"
	"encoding/json"
	"log"

	"github.com/nlopes/slack/slackevents"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func Record(ctx context.Context, m PubSubMessage) error {
	event := &slackevents.MessageEvent{}
	err := json.Unmarshal(m.Data, event)
	if err != nil {
		return err
	}
	log.Printf("%+v", event)
	return nil
}
