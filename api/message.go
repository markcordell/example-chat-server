package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/markcordell/example-chat-server/internal/datastore"
	"github.com/pkg/errors"
)

func (a *Api) PostMessage(c http.ResponseWriter, req *http.Request) {
	var i datastore.IncomingMessage

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorReturn := errors.Wrap(err, "error reading body")
		log.Print(errorReturn)
		c.WriteHeader(http.StatusInternalServerError)
		c.Write([]byte(errorReturn.Error()))
		return
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		errorReturn := errors.Wrap(err, "json error")
		log.Print(errorReturn)
		c.WriteHeader(http.StatusBadRequest)
		c.Write([]byte(errorReturn.Error()))
		return
	}

	err = validateData(i.User, i.Text)
	if err != nil {
		errorReturn := errors.Wrap(err, "error validating data")
		log.Print(errorReturn)
		c.WriteHeader(http.StatusBadRequest)
		c.Write([]byte(errorReturn.Error()))
		return
	}

	err = a.queue.PutMessageIntoQueue(i.User, i.Text)
	if err != nil {
		errorReturn := errors.Wrap(err, "error placing message into queue")
		log.Print(errorReturn)
		c.WriteHeader(http.StatusInternalServerError)
		c.Write([]byte(errorReturn.Error()))
		return
	}
	err = a.users.CheckAndAddNewUser(i.User)
	if err != nil {
		errorReturn := errors.Wrap(err, "error placing user into queue")
		log.Print(errorReturn)
		c.WriteHeader(http.StatusInternalServerError)
		c.Write([]byte(errorReturn.Error()))
		return
	}
	c.Write([]byte("{\"ok\": true}"))

}

func validateData(user string, text string) error {
	if len(user) == 0 {
		return errors.New("User must be provided")
	}

	if len(text) == 0 {
		return errors.New("Message text must be provided.")
	}
	return nil
}
