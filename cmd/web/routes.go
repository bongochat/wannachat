package main

import (
	"net/http"

	"github.com/bongochat/wannachat/pkg/config"
	"github.com/bongochat/wannachat/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// use middleware
	mux.Use(middleware.Recoverer)
	// custom middleware
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomePageHandler)
	mux.Get("/faq/", handlers.Repo.FAQPageHandler)
	mux.Get("/tos/", handlers.Repo.TOSPageHandler)
	mux.Get("/safety-tips/", handlers.Repo.SaftyTipsPageHandler)
	mux.Get("/chat-rules/", handlers.Repo.ChatRulesPageHandler)

	mux.Get("/login/", handlers.Repo.LoginPageHandler)
	mux.Get("/signup/", handlers.Repo.SignupPageHandler)

	mux.Get("/contact/", handlers.Repo.ContactPageHandler)
	mux.Post("/contact/", handlers.Repo.PostContactPageHandler)
	mux.Get("/contact/json/", handlers.Repo.ContactPageJSON)

	mux.Get("/robots.txt", handlers.Repo.RobotsHandler)

	// load static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
