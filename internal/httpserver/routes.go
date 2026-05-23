package httpserver

import (
	ah "github.com/eugbyte/sqlc_tutorial/internal/api/author/handler"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *ah.AuthorHandler) {
	r.Route("/authors", func(r chi.Router) {
		r.Post("/", h.CreateAuthor)
		r.Get("/", h.ListAuthors)
		r.Get("/{authorID}", h.GetAuthor)
		r.Put("/{authorID}", h.UpdateAuthor)
		r.Delete("/{authorID}", h.DeleteAuthor)
	})
}
