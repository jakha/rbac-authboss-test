package controllers

import (
	"net/http"
)

func Bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar"))
	w.WriteHeader(200)
}
