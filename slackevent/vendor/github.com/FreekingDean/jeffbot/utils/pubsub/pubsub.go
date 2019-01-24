package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
)

var client *pubsub.Client

func init() {
	var err error
	client, err = pubsub.NewClient(context.Background(), "jeffbot")
	if err != nil {
		panic(err)
	}
}

func Publish(ctx context.Context, topic string, data []byte) error {
	res := client.Topic(topic).Publish(ctx, &pubsub.Message{Data: data})
	_, err := res.Get(ctx)
	return err
}
