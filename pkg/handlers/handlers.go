package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bongochat/wannachat/pkg/config"
	"github.com/bongochat/wannachat/pkg/models"
	"github.com/bongochat/wannachat/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.go.tmpl", &models.TemplateData{})
}

func (m *Repository) FAQPageHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "faq.page.go.tmpl", &models.TemplateData{})
}

func (m *Repository) BlogPageHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "blog.page.go.tmpl", &models.TemplateData{})
}

func (m *Repository) ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.go.tmpl", &models.TemplateData{})
}

func (m *Repository) PostContactPageHandler(w http.ResponseWriter, r *http.Request) {
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	phoneNumber := r.Form.Get("phoneNumber")
	message := r.Form.Get("message")
	w.Write([]byte(fmt.Sprintf("Name: %s, Email: %s, Phone Number: %s, Message: %s", name, email, phoneNumber, message)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) ContactPageJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Thanks for sending us message. We will contact you soon",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) RobotsHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "robots.page.go.tmpl", &models.TemplateData{})
}
