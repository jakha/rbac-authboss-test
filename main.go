package main

import (
	"github.com/jakha/rbac-authboss-test/config"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", config.StartApp())
}
