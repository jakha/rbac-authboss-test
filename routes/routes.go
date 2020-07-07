package routes

import (
	"github.com/jakha/rbac-authboss-test/controllers"
	"net/http"
)

type NextHandler func(http.Handler) http.Handler

type Routing struct {
	H        http.Handler
	MidChain []NextHandler
}

func Init() {
	http.HandleFunc("/", controllers.Main)
	http.HandleFunc("/auth", controllers.Auth)
	http.HandleFunc("/foo", controllers.Foo)
	http.HandleFunc("/bar", controllers.Bar)
	http.HandleFunc("/sigma", controllers.Sigma)
}

func (rout *Routing) HookUp(h http.Handler) {
	midChainLen := len(rout.MidChain)

	for i := midChainLen - 1; i > -1; i-- {
		h = rout.MidChain[i](h)
	}

	rout.H = h
}

func (rout Routing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rout.H.ServeHTTP(w, r)
	return
}
