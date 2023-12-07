package main

import (
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

// write message to console in every url hit
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit the page", r.URL.String())
		next.ServeHTTP(w, r)
	})
}

// adds CSRF protection to all post requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// load and save the session in every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
