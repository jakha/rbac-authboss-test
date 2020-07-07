package main

import (
	"github.com/go-chi/chi"
	authboss_components "github.com/jakha/rbac-authboss-test/authboss-components"
	"github.com/jakha/rbac-authboss-test/routes"
	"github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/volatiletech/authboss/v3/remember"

	//"github.com/volatiletech/authboss/v3/remember"
	"net/http"
)

func main() {
	//setUp()
	httpserv()
}

type simple struct {
}

func (s simple) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("third"))
}

func httpserv() {
	routes.Init()

	var rr = routes.Routing{}

	rr.MidChain = append(rr.MidChain, ReturnHandle, ReturnSecondHandle)
	rr.HookUp(simple{})

	http.ListenAndServe(":8080", rr)
}

func setUp() {
	ab := authboss.New()
	mux := chi.NewRouter()
	ab.Config.Storage.Server = authboss_components.ServerStorage{}
	ab.Config.Storage.SessionState = authboss_components.ClientStateReadWriter{}
	ab.Config.Storage.CookieState = authboss_components.ClientStateReadWriter{}

	mux.Use(ab.LoadClientStateMiddleware, remember.Middleware(ab))

	ab.Config.Paths.Mount = "/td"
	ab.Config.Paths.RootURL = "http://localhost:8080"

	// This is using the renderer from: github.com/volatiletech/authboss
	ab.Config.Core.ViewRenderer = abrenderer.NewHTML("/auth", "ab_views")
	// Probably want a MailRenderer here too.

	// Set up defaults for basically everything besides the ViewRenderer/MailRenderer in the HTTP stack
	defaults.SetCore(&ab.Config, false, false)

	if err := ab.Init(); err != nil {
		panic(err)
	}

	//// Authed routes
	//mux.Group(func(mux chi.Router) {
	//mux.Use(authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondUnauthorized), lock.Middleware(ab), confirm.Middleware(ab))
	//	mux.MethodFunc("GET", "/blogs/new", newblog)
	//	mux.MethodFunc("GET", "/blogs/{id}/edit", edit)
	//	mux.MethodFunc("POST", "/blogs/{id}/edit", update)
	//	mux.MethodFunc("POST", "/blogs/new", create)
	//	// This should actually be a DELETE but can't be bothered to make a proper
	//	// destroy link using javascript atm. See where AB allows you to configure
	//	// the logout HTTP method.
	//	mux.MethodFunc("GET", "/blogs/{id}/destroy", destroy)
	//})

	// Mount the router to a path (this should be the same as the Mount path above)
	// mux in this example is a chi router, but it could be anything that can route to
	// the Core.Router.
	mux.Mount("/authboss", http.StripPrefix("/authboss", ab.Config.Core.Router))

	http.ListenAndServe(":8080", mux)
}

func ReturnHandle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("first"))
		h.ServeHTTP(w, r)
	})
}

func ReturnSecondHandle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("second"))
		h.ServeHTTP(w, r)
	})
}
