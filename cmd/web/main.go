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

//setting global variables
const port = ":8090"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	//configuration of sessions
	session = scs.New()
	session.Lifetime = 12 * time.Hour // session should last for half aday
	//store session into Cookie
	session.Cookie.Persist = true //ensures your session persist even if you close your web browser
	session.Cookie.Secure = app.InProduction // encrypt the cookie
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session


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
