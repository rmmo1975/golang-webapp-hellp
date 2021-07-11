package handlers

import (
	"net/http"

	"example.com/pkg/config"
	"example.com/pkg/models"
	"example.com/pkg/render"
)

// used by tthe handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
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
	render.RenderTemplate(w, "home-page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// add some logic here....
	stringMap := make(map[string]string)
	stringMap["test"] = "bolsonaro gosta de cheirar virilha suja"

	render.RenderTemplate(w, "about-page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
