package messages

type Message struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Raw    []byte `json:"raw"`
}

type ResponseRequest struct {
	Original         string      `json:"original"`
	Response         string      `json:"response"`
	ResponseSource   string      `json:"response_source"`
	ResponseMetadata interface{} `json:"response_metadata"`
}
