package handlers

import (
	"net/http"

	"github.com/sanyog2491/bookings/package/config"
	"github.com/sanyog2491/bookings/package/model"
	"github.com/sanyog2491/bookings/package/render"
)

var Repo *Repository

type Repository struct {
	App *config.Appconfig
}

func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home3(w http.ResponseWriter, r *http.Request) {

	/*everytime somebody hits the hhome page for that users session it
	will store the IP as a string in remote session*/
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Rendertemplate(w, "home.page.tmpl", &model.TemplateData{})
}
func (m *Repository) About2(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello.again."

	// now we will get that Remote Ip and we will sto0re it inside the map
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	//send data to the template
	render.Rendertemplate(w, "about.page.tmpl", &model.TemplateData{
		StringMap: stringMap,
	})
}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.Rendertemplate(w, "make.reservation.tmpl", &model.TemplateData{})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.Rendertemplate(w, "generals.page.tmpl", &model.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.Rendertemplate(w, "majors.page.tmpl", &model.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Rendertemplate(w, "search-availability.page.tmpl", &model.TemplateData{})
}
