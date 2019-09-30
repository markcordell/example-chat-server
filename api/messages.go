package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/markcordell/example-chat-server/internal/queue"
)

func (a *Api) GetMessages(c http.ResponseWriter, req *http.Request) {
	outputQueue := new(queue.Queue)

	if len(a.queue.Messages) >= 100 {
		outputQueue.Messages = a.queue.Messages[len(a.queue.Messages)-100 : len(a.queue.Messages)]
	} else {
		outputQueue.Messages = a.queue.Messages
	}

	output, err := json.Marshal(outputQueue)
	if err != nil {
		errorReturn := fmt.Sprintf("error decoding json of queue", err)
		log.Print(errorReturn)
		c.WriteHeader(http.StatusInternalServerError)
		c.Write([]byte(errorReturn))
	}
	c.Write([]byte(output))
}
