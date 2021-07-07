package handlers

import (
	"net/http"

	"example.com/pkg/config"
	"example.com/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func CreateRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func StartHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home-page.html")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about-page.html")
}
