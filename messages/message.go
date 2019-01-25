package messages

import (
	"context"
	"encoding/json"
)

type Message struct {
	Text   string `json:"text" bigquery:"text"`
	Source string `json:"source" bigquery:"source"`
	Raw    []byte `json:"raw" bigquery:"raw"`
}

type ResponseRequest struct {
	Original         string      `json:"original"`
	Response         string      `json:"response"`
	ResponseSource   string      `json:"response_source"`
	ResponseMetadata interface{} `json:"response_metadata"`
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func Parse(ctx context.Context, m PubSubMessage) error {
	message := &Message{}
	err := json.Unmarshal(m.Data, message)
	if err != nil {
		return err
	}
	return table.Inserter().Put(ctx, message)
}
