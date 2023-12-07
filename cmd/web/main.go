package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bongochat/wannachat/pkg/config"
	"github.com/bongochat/wannachat/pkg/handlers"
	"github.com/bongochat/wannachat/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	// create session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = app.InProduction

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on http://127.0.0.1" + portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
