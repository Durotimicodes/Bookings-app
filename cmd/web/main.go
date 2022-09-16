package main

import (
	"fmt"
	"net/http"

	"github.com/theclassicdev/monolithic-app/pkg/handlers"
)

const port = ":8089"

func main() {

	//handlers
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	//server
	fmt.Printf("Starting server on port %s", port)
	_ = http.ListenAndServe(port, nil)
}
