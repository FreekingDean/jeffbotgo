package messages

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"gopkg.in/jdkato/prose.v2"
	"log"
	"math/rand"
)

const baseQuery1 = `
 WITH message as (SELECT n_grams as grams FROM messages.messages WHERE ARRAY_LENGTH(n_grams) > 0) SELECT gram_1, gram_2, gram_3, count(*) as popularity FROM message, unnest(grams) WHERE gram_1 = "%s" GROUP BY gram_1, gram_2, gram_3 ORDER BY popularity DESC LIMIT 10
 `

const baseQuery2 = `
 WITH message as (SELECT n_grams as grams FROM messages.messages WHERE ARRAY_LENGTH(n_grams) > 0) SELECT gram_1, gram_2, gram_3, count(*) as popularity FROM message, unnest(grams) WHERE gram_1 = "%s" AND gram_2 = "%s" GROUP BY gram_1, gram_2, gram_3 ORDER BY popularity DESC LIMIT 10
 `

type ResponseRequest struct {
	Original         string      `json:"original"`
	Response         string      `json:"response"`
	ResponseSource   string      `json:"response_source"`
	ResponseMetadata interface{} `json:"response_metadata"`
}

type Grams struct {
	NGram
	Popularity int `bigquery:"popularity"`
}

func GenerateResponse(ctx context.Context, m PubSubMessage) error {
	decoded, err := base64.StdEncoding.DecodeString(string(m.Data))
	if err == nil {
		m.Data = decoded
	}
	req := &ResponseRequest{}
	err = json.Unmarshal(m.Data, req)
	if err != nil {
		return err
	}

	doc, err := prose.NewDocument(req.Original)
	if err != nil {
		return err
	}
	subject := doc.Tokens()[0].Text
	log.Println(subject)
	for _, tkn := range doc.Tokens() {
		if tkn.Tag == "NN" || tkn.Tag == "NNP" || tkn.Tag == "NNPS" || tkn.Tag == "NNS" {
			subject = tkn.Text
			log.Println(subject)
			break
		}
	}

	q := client.Query(fmt.Sprintf(baseQuery1, subject))
	it, err := q.Read(ctx)
	if err != nil {
		return err
	}

	var grams Grams
	err = it.Next(&grams)
	if err != nil {
		return err
	}
	sentence := []string{subject, grams.Gram2}

	g1 := subject
	g2 := grams.Gram2
	for {
		if len(sentence) > 20 {
			break
		}
		q = client.Query(fmt.Sprintf(baseQuery2, g1, g2))
		it, err = q.Read(ctx)
		if err != nil {
			return err
		}
		if it.TotalRows <= 0 {
			break
		}
		for {
			err = it.Next(&grams)
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}
			if rand.Float32() < 0.85 {
				break
			}
		}
		sentence = append(sentence, grams.Gram3)
	}
	log.Println(sentence)
	return nil
}
