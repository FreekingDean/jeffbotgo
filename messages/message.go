package messages

type Message struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Raw    []byte `json:"raw"`
}
