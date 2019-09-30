package api

import (
	"net/http"
)

func (a *Api) Status(c http.ResponseWriter, req *http.Request) {
	c.Write([]byte("alive"))
}
