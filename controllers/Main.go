package controllers

import (
	"net/http"
)

func Main(w http.ResponseWriter, r *http.Request) {
	//http.Redirect()
	w.WriteHeader(200)
}
