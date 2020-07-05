package routes

import (
	"github.com/jakha/rbac-authboss-test/controllers"
	"net/http"
)

func Init() {
	http.HandleFunc("/", controllers.Main)
	http.HandleFunc("/auth", controllers.Auth)
	http.HandleFunc("/foo", controllers.Foo)
	http.HandleFunc("/bar", controllers.Bar)
	http.HandleFunc("/sigma", controllers.Sigma)
}

