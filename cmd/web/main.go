package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/theclassicdev/monolithic-app/pkg/config"
	"github.com/theclassicdev/monolithic-app/pkg/handlers"
	"github.com/theclassicdev/monolithic-app/pkg/render"
)

const port = ":8089"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// store template cache into the variable app
	app.TemplateCache = tc

	//handlers
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	//server
	fmt.Printf("Starting server on port %s", port)
	_ = http.ListenAndServe(port, nil)
}
