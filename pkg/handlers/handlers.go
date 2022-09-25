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

func (R *Repository) Home(w http.ResponseWriter, r *http.Request) {

	/*everytime a client makes a request to this page this will store the
	clients Ip address (IPv4 or IPv6)*/
	remoteIP := r.RemoteAddr
	R.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})

}

func (R *Repository) About(w http.ResponseWriter, r *http.Request) {
	sm := make(map[string]string)
	sm["test"] = "Hello monolithic application"

	//grab the session value gotten from the HomePage
	remoteIP := R.App.Session.GetString(r.Context(), "remote_ip")
	sm["remote_ip"] = remoteIP

	//send data to template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: sm,
	})
}


//Reservation renders the make a reservation page and displays a form
func (R *Repository) Reservation(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "make-reservation.html", &models.TemplateData{})
}

//General renders the generals quarters page
func (R *Repository) Generals(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "generals-quarters.html", &models.TemplateData{})
}
//Majors renders the room page
func (R *Repository) Majors(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "majors-suite.html", &models.TemplateData{})
}


//Availability renders the available room page
func (R *Repository) Availability(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "search-availability.html", &models.TemplateData{})
}