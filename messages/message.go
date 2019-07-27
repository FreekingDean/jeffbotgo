package messages

import (
	"strings"
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
)

type Message struct {
	Text   string  `json:"text" bigquery:"text"`
	NGrams []NGram `json:"n_gram" bigquery:"n_gram"`
	Source string  `json:"source" bigquery:"source"`
	Raw    []byte  `json:"raw" bigquery:"raw"`
}

type NGram struct {
	Gram1 string `json:"gram_1", bigquery:gram_1"`
	Gram2 string `json:"gram_2", bigquery:gram_2"`
	Gram3 string `json:"gram_3", bigquery:gram_3"`
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
	decoded, err := base64.StdEncoding.DecodeString(string(m.Data))
	log.Println(string(m.Data))
	log.Println(string(decoded))
	if err == nil {
		m.Data = decoded
	}
	log.Println(string(m.Data))

	message := &Message{}
	err = json.Unmarshal(m.Data, message)
	if err != nil {
		return err
	}
	message.NGrams = []NGram{}
	word1 := ""
	word2 := ""
	word3 := ""
	for i, word := strings.Split(message.Text, " ") {
		word1 = word2
		word2 = word3
		word3 = word
		message.NGrams = append(message.NGrams, NGram{word1, word2, word3})
	}
	return table.Inserter().Put(ctx, message)
}
