package slackmessagerecorder

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	"github.com/nlopes/slack/slackevents"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Message struct {
	Source  string `firestore:"source"`
	Message string `firestore:"message"`
	Raw     string `firestore:"raw"`
}

// HelloPubSub consumes a Pub/Sub message.
func Record(ctx context.Context, m PubSubMessage) error {
	event := &slackevents.MessageEvent{}
	err := json.Unmarshal(m.Data, event)
	if err != nil {
		return err
	}
	client, err := firestore.NewClient(ctx, "jeffbot")
	if err != nil {
		return err
	}

	_, _, err = client.Collection("messages").Add(ctx, &Message{
		Source:  "slack",
		Message: event.Text,
		Raw:     string(m.Data),
	})
	return err
}
