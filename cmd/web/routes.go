package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/theclassicdev/monolithic-app/pkg/config"
	"github.com/theclassicdev/monolithic-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//create a multiplexer
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) //Recovers from panic , logs the panic and return an http 500
	mux.Use(WriteToConsole)       //customized middelware
	mux.Use(NoSurf)               //built to create a CSRFToken
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomePage)
	mux.Get("/about", handlers.Repo.AboutPage)

	//serving the static file
	//http.FileServer return the handler that the http request which hold the content
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
