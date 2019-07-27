package messages

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"

	"cloud.google.com/go/bigquery"
)

type Message struct {
	Text   string  `json:"text" bigquery:"text"`
	NGrams []NGram `json:"n_grams" bigquery:"n_gram"`
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
	words := strings.Split(message.Text, " ")
	for i := 0; i < len(words); i++ {
		word1 := words[i]
		word2 := ""
		if i < len(words)-1 {
			word2 = words[i+1]
		}
		word3 := ""
		if i < len(words)-2 {
			word3 = words[i+2]
		}
		message.NGrams = append(message.NGrams, NGram{word1, word2, word3})
	}
	errs := table.Inserter().Put(ctx, message)
	if errs != nil {
		log.Println(err.(bigquery.PutMultiError)[0].Error())
		return err
	}
	return nil
}
