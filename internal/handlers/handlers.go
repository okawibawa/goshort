package handlers

import (
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	store     *pgxpool.Pool
	templates *template.Template
}

func NewHandler(db *pgxpool.Pool) *Handler {
	templates := template.Must(template.ParseGlob("web/templates/*.html"))
	return &Handler{store: db, templates: templates}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "index.html", nil)
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
}
