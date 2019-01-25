package messages

import (
	"context"

	"cloud.google.com/go/bigquery"
)

var table *bigquery.Table

func init() {
	client, err := bigquery.NewClient(context.Background(), "jeffbot")
	if err != nil {
		panic(err)
	}
	table = client.Dataset("messages").Table("messages")
}
