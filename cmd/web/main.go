package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/theclassicdev/monolithic-app/pkg/config"
	"github.com/theclassicdev/monolithic-app/pkg/handlers"
	"github.com/theclassicdev/monolithic-app/pkg/render"
)

const port = ":8089"

func main() {

	var app config.AppConfig

	//configuration of sessions
	session := scs.New()
	session.Lifetime = 24 * time.Hour // session shoulf last for a day

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// store template cache into the variable app
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	//server
	fmt.Printf("Starting server on port %s", port)
	// _ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
