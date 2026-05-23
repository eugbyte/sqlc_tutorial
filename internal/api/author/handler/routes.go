package handler

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *AuthorHandler) {
	r.Route("/authors", func(r chi.Router) {
		r.Post("/", h.CreateAuthor)
		r.Get("/", h.ListAuthors)
		r.Get("/{authorID}", h.GetAuthor)
		r.Put("/{authorID}", h.UpdateAuthor)
		r.Delete("/{authorID}", h.DeleteAuthor)
	})
}
