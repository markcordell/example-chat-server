package queue

import (
	"time"

	"github.com/markcordell/example-chat-server/internal/datastore"
)

type Queue struct {
	Messages []datastore.Message `json:"messages"` // A slice was used instead of an array because appending to and poping from an array is easier.
}

func (q *Queue) PutMessageIntoQueue(user string, text string) error {
	timeStamp := time.Now().Unix()
	message := datastore.Message{
		TimeStamp: timeStamp,
		User:      user,
		Text:      text,
	}

	q.Messages = append(q.Messages, message)
	return nil
}
