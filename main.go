package main

import (
	"github.com/jakha/rbac-authboss-test/routes"
	"net/http"
)

func main() {
	routes.Init()
	http.ListenAndServe(":8080", nil)
}
