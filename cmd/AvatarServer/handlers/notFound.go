package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/context"
)

//NotFound handles all 404 not found requests
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	context.Set(r, errorKey, errors.New("Endpoint does not exist: "+r.RequestURI))
	Error(w, r)
}
