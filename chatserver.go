package main

import (
	"net/http"

	"github.com/markcordell/example-chat-server/api"

	"goji.io"
	"goji.io/pat"
)

func main() {
	a := api.ConstructApi()
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/status"), a.Status)

	mux.HandleFunc(pat.Get("/messages"), a.GetMessages)

	mux.HandleFunc(pat.Post("/message"), a.PostMessage)

	mux.HandleFunc(pat.Get("/users"), a.GetUsers)

	err := http.ListenAndServe(":8081", mux)
	panic(err)
}
