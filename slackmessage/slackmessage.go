package slackmessage

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
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
func Parse(ctx context.Context, m PubSubMessage) error {
	decoded, err := base64.StdEncoding.DecodeString(string(m.Data))
	log.Println(string(m.Data))
	log.Println(string(decoded))
	if err == nil {
		m.Data = decoded
	}
	log.Println(string(m.Data))

	event := &slackevents.Message{}
	err = json.Unmarshal(m.Data, event)
	if err != nil {
		return err
	}
	log.Printf("%+v\n", event)

	if strings.Contains(event.Text, "jeffbot!!") {
		req := messages.ResponseRequest{
			Original:       event.Text,
			ResponseSource: "slack",
			ResponseMetadata: SlackResponseRequest{
				Channel: event.Channel,
			},
		}
		err = pubsub.Publish(ctx, "jeffbot.request_response", req)
		if err != nil {
			return err
		}
	}

	message := &messages.Message{
		Text:   event.Text,
		Source: "slack",
		Raw:    m.Data,
	}

	return pubsub.Publish(ctx, "jeffbot.message", message)
}
