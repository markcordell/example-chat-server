package datastore

type IncomingMessage struct {
	User string
	Text string
}

type Message struct {
	TimeStamp int64  `json:"timestamp"` // timestamp is a unix timestamp which is represented as a in64
	User      string `json:"user"`
	Text      string `json:"text"`
}
