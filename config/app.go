package config

import (
	"encoding/base64"
	"github.com/jakha/rbac-authboss-test/routes"
	"github.com/justinas/nosurf"
	abclientstate "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"log"
	"net/http"
)

const (
	sessionCookieName = "ab_test"
)

var (
	ab *authboss.Authboss = authboss.New()
)

func StartApp() http.Handler {
	routes.InitRoutes()
	return initMids()
}

func initMids() http.Handler {
	prepareAuthboss()
	h := ab.LoadClientStateMiddleware(http.DefaultServeMux)
	h = nosurfing(h)
	h = authboss.Middleware2(ab, authboss.RequireNone, authboss.RespondRedirect)(h)
	//h = lock.Middleware(ab)(h)
	//h = confirm.Middleware(ab)(h)

	return h
}

func prepareAuthboss() {
	ab.Config.Paths.Mount = "/auth"
	ab.Config.Paths.RootURL = "http://localhost:8080/"
	ab.Config.Storage.SessionState = wantSessionStorer()
	ab.Config.Storage.CookieState = wantCookieStorer()
	ab.Config.Storage.Server = GetDb()
}

func wantSessionStorer() authboss.ClientStateReadWriter {
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`1/65+JA==`)
	return abclientstate.NewSessionStorer(sessionCookieName, sessionStoreKey, nil)
}

func wantCookieStorer() authboss.ClientStateReadWriter {
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`1+2==`)
	cookieStore := abclientstate.NewCookieStorer(cookieStoreKey, nil)
	cookieStore.HTTPOnly = false
	cookieStore.Secure = false
	return cookieStore
}

func nosurfing(h http.Handler) http.Handler {
	surfing := nosurf.New(h)
	surfing.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Failed to validate CSRF token:", nosurf.Reason(r))
		w.WriteHeader(http.StatusBadRequest)
	}))
	return surfing
}
