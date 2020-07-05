package controllers

import (
	"github.com/jakha/rbac-authboss-test/helpers"
	"io/ioutil"
	"net/http"
	"os"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		handleAuthForm(w)
	case http.MethodPost:

	}

}

func handleAuthForm(w http.ResponseWriter) {
	file, err := os.Open("./static/auth.html")
	helpers.Check(err)
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	helpers.Check(err)

	w.Write(data)
}
