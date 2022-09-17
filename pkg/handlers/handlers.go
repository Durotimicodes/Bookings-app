package handlers

import (
	"net/http"

	"github.com/theclassicdev/monolithic-app/pkg/config"
	"github.com/theclassicdev/monolithic-app/pkg/models"
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
	render.RenderTemplate(w, "home.html", &models.TemplateData{})

}

func (R *Repository) AboutPage(w http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)

	sm["test"] = "Hello monolithic application"

	//send data to template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: sm,
	})
}
