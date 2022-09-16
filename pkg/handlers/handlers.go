package handlers

import (
	"net/http"

	"github.com/theclassicdev/monolithic-app/pkg/config"
	"github.com/theclassicdev/monolithic-app/pkg/render"
)

//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (R *Repository) HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")

}

func (R *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
