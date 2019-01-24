package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
)

var client *pubsub.Client

func init() {
	var err error
	client, err = pubsub.NewClient(context.Background(), "jeffbot")
	if err != nil {
		panic(err)
	}
}

func Publish(ctx context.Context, topic string, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res := client.Topic(topic).Publish(ctx, &pubsub.Message{Data: dataBytes})
	_, err = res.Get(ctx)
	return err
}
