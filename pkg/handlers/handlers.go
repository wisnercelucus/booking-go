package handlers

import (
	"net/http"

	"github.com/wisnercelucus/udemy-booking/pkg/config"
	"github.com/wisnercelucus/udemy-booking/pkg/models"
	"github.com/wisnercelucus/udemy-booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	templ := "home.page.html"
	render.RenderTemplate(w, templ, &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"
	templ := "about.page.html"
	render.RenderTemplate(w, templ, &models.TemplateData{
		StringMap: stringMap,
	})
}
