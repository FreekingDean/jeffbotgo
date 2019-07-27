package messages

import (
	"context"

	"cloud.google.com/go/bigquery"
)

var table *bigquery.Table
var client *bigquery.Client

func init() {
	var err error
	client, err = bigquery.NewClient(context.Background(), "jeffbot")
	if err != nil {
		panic(err)
	}
	table = client.Dataset("messages").Table("messages")
}
