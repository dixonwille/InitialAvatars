package handlers

import (
	"net/http"

	"github.com/dixonwille/InitialAvatars"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type key int

const (
	errorKey key = iota
)

//Avatar handles all request to create an Avatar
func Avatar(w http.ResponseWriter, r *http.Request) {
	initials := mux.Vars(r)["initials"]
	colorHex := r.URL.Query().Get("color")
	var color *avatar.Color
	var err error

	//create color for the avatar
	if colorHex == "" {
		color = avatar.RandomColor()
	} else {
		color, err = avatar.ColorFromHex(colorHex)
		if err != nil {
			if avatar.IsInvalidValue(err) || avatar.IsValueOutOfRange(err) {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			context.Set(r, errorKey, err)
			Error(w, r)
			return
		}
	}

	//Creqte the avatar and return it to the browser
	myAvatar, err := avatar.NewAvatar(initials, *color)
	if err != nil {
		if avatar.IsInvalidValue(err) || avatar.IsValueOutOfRange(err) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		context.Set(r, errorKey, err)
		Error(w, r)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	myAvatar.Create(w)
}
