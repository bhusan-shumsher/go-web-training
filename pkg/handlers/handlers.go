package handlers

import (
	"net/http"

	"github.com/bhusan-shumsher/go-web-training/pkg/config"
	"github.com/bhusan-shumsher/go-web-training/pkg/render"
)

var Repo *Respository

type Respository struct {
	App *config.AppConfig
}

// creates the repository
func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

// sets the respository for the handlers
func NewHandlers(r *Respository) {
	Repo = r
}
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
