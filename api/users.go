package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (a *Api) GetUsers(c http.ResponseWriter, req *http.Request) {
	output, err := json.Marshal(a.users)
	if err != nil {
		errorReturn := fmt.Sprintf("error getting users", err)
		log.Print(errorReturn)
		c.WriteHeader(http.StatusInternalServerError)
		c.Write([]byte(errorReturn))
		return
	}
	c.Write([]byte(output))

}
