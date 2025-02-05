package handlers

import (
	"net/http"
	"strconv"

	"github.com/rainclear/accroo/pkg/config"
	"github.com/rainclear/accroo/pkg/dbm"
	"github.com/rainclear/accroo/pkg/models"
	"github.com/rainclear/accroo/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}

func (m *Repository) AccountTypes(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	for i, account_type := range Repo.App.AccountTypes {
		index := strconv.Itoa(i)
		stringMap[index] = account_type
	}

	render.RenderTemplate(w, "account_types.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) AccountCategories(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	account_categories, _ := dbm.ListAccountCategories()

	for i, account_category := range account_categories {
		index := strconv.Itoa(i)
		stringMap[index] = account_category
	}

	render.RenderTemplate(w, "account_categories.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) ModifyAccount(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "modify_account.page.html", &models.TemplateData{})
}
