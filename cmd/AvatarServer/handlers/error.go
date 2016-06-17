package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
)

type errorMessage struct {
	Error string `json:"error"`
}

func newErrorMessage(msg string) *errorMessage {
	return &errorMessage{
		Error: msg,
	}
}

//Error wirtes and error in json to writer
func Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err, ok := context.Get(r, errorKey).(error)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		//42 is the code because that is the answer to life
		w.Write([]byte("Something went wrong. Contact maintainer. Code: 40"))
		return
	}
	msg := newErrorMessage(err.Error())
	content, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong. Contact maintainer. Code: 41"))
		return
	}
	w.Write(content)
}
