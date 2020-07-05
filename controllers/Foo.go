package controllers

import "net/http"

func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Foo"))
	w.WriteHeader(200)
}