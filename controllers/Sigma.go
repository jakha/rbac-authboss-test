package controllers

import "net/http"

func Sigma(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sigma"))
	w.WriteHeader(200)
}
