package controllers

import (
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
