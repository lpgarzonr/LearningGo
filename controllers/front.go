package controllers

import "net/http"

func RegisterControllers() {
	sc := newSongController()
	http.Handle("/songs", *sc)
	http.Handle("/songs/", *sc)
}
