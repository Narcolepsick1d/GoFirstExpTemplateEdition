package handlers

import (
	"awesomeProject2/pkg/config"
	"awesomeProject2/pkg/models"
	"awesomeProject2/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	StringMap := make(map[string]string)
	StringMap["test"] = "hello again"

	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{
		StringMap: StringMap,
	})

}
