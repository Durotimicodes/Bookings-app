package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/theclassicdev/monolithic-app/internal-pkg/config"
	"github.com/theclassicdev/monolithic-app/internal-pkg/models"
	"github.com/theclassicdev/monolithic-app/internal-pkg/render"
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

	render.RenderTemplate(w, "home.html", &models.TemplateData{}, r)

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
	}, r)
}

//Reservation renders the make a reservation page and displays a form
func (R *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "make-reservation.html", &models.TemplateData{}, r)
}

//General renders the generals quarters page
func (R *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "general.html", &models.TemplateData{}, r)
}

//Majors renders the room page
func (R *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.html", &models.TemplateData{}, r)
}

//Availability renders the available room page
func (R *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search-availability.html", &models.TemplateData{}, r)
}

//structure of json response
type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//Availability renders the available room page
func (R *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

//AvailabilityJSON handles request for availability and returns a json
func (R *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-type", "application/json")

	w.Write(out)

}

func (R *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.html", &models.TemplateData{}, r)
}
