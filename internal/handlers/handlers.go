package handlers

import (
	"context"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/okawibawa/goshort/internal/shortener"
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

func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	shortCode, err := shortener.GenerateShortCode()
	if err != nil {
		http.Error(w, "Error generating short code.", http.StatusInternalServerError)
		return
	}

	_, err = h.store.Exec(context.Background(), "insert into urls (original_url, shorten_url) values ($1, $2)", url, shortCode)
	if err != nil {
		http.Error(w, "Error generating short code.", http.StatusInternalServerError)
		return
	}

	data := struct {
		OriginalURL  string
		ShortenedURL string
	}{
		OriginalURL:  url,
		ShortenedURL: "https://www.goshorty.okawibawa.dev/" + shortCode,
	}

	h.templates.ExecuteTemplate(w, "result.html", data)
}
