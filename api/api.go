package api

import (
	"github.com/markcordell/example-chat-server/internal/datastore"
	"github.com/markcordell/example-chat-server/internal/queue"
)

type Api struct {
	queue queue.Queue
	users datastore.Users
}

func ConstructApi() *Api {
	var a Api
	a.users = *datastore.ConstructUsers()
	return &a
}
