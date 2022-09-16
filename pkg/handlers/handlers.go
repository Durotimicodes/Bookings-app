package handlers

import (
	"net/http"

	"github.com/theclassicdev/monolithic-app/pkg/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")

}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
