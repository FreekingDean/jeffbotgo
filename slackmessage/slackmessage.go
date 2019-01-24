package slackmessage

import (
	"context"
	"encoding/json"

	"github.com/FreekingDean/jeffbotgo/messages"
	"github.com/FreekingDean/jeffbotgo/utils/pubsub"
	"github.com/FreekingDean/slackevents"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func Record(ctx context.Context, m PubSubMessage) error {
	event := &slackevents.Message{}
	err := json.Unmarshal(m.Data, event)
	if err != nil {
		return err
	}
	&messages.Message{
		Text:   event.Text,
		Source: "slack",
		Raw:    event,
	}
	pubsub.Publish(ctx, "jeffbot.message")
}
