package slackmessage

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/FreekingDean/jeffbotgo/messages"
	"github.com/FreekingDean/jeffbotgo/utils/pubsub"
	"github.com/FreekingDean/slackevents"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type SlackResponseRequest struct {
	Channel string `json:"channel"`
}

// HelloPubSub consumes a Pub/Sub message.
func Record(ctx context.Context, m PubSubMessage) error {
	event := &slackevents.Message{}
	err := json.Unmarshal(m.Data, event)
	if err != nil {
		return err
	}

	if strings.Contains(event.Text, "jeffbot!!") {
		req := messages.ResponseRequest{
			Original:       event.Text,
			ResponseSource: "slack",
			ResponseMetadata: SlackResponseRequest{
				Channel: event.Channel,
			},
		}
		err = pubsub.Publish(ctx, "jeffbot.request_response", req)
	}

	message := &messages.Message{
		Text:   event.Text,
		Source: "slack",
		Raw:    m.Data,
	}

	return pubsub.Publish(ctx, "jeffbot.message", message)
}
