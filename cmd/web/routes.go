package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/theclassicdev/monolithic-app/internal-pkg/config"
	"github.com/theclassicdev/monolithic-app/internal-pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	//create a multiplexer
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) //Recovers from panic , logs the panic and return an http 500
	mux.Use(WriteToConsole)       //customized middelware
	mux.Use(NoSurf)               //built to create a CSRFToken
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)


	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)

	//serving the static file
	//http.FileServer return the handler that the http request which hold the content
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
