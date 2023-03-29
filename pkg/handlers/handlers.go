package handlers

import (
	"net/http"

	"github.com/bhusan-shumsher/go-web-training/pkg/config"
	"github.com/bhusan-shumsher/go-web-training/pkg/models"
	"github.com/bhusan-shumsher/go-web-training/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// creates the repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the Repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}
