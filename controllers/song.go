package controllers

import (
	"net/http"
	"regexp"
)

type songContoller struct {
	songIdPattern *regexp.Regexp
}

// method ServerHTTP bind to song controller getting request
func (sc songContoller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Song Controller"))
}

func newSongController() *songContoller {
	return &songContoller{
		songIdPattern: regexp.MustCompile(`^/songs/(\d+)/?`),
	}
}
