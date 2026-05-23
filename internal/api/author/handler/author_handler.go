package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/eugbyte/sqlc_tutorial/internal/api/author/service"
	"github.com/eugbyte/sqlc_tutorial/internal/domain/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
)

type AuthorHandler struct {
	authorService *service.AuthorService
}

func NewAuthorHandler(authorService *service.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var payload model.AuthorRequest
	if err := decodeJSON(r, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	author, err := h.authorService.CreateAuthor(r.Context(), payload)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to create author"})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, author)
}

func (h *AuthorHandler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, ok := parseAuthorID(w, r)
	if !ok {
		return
	}

	author, err := h.authorService.GetAuthor(r.Context(), authorID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": "author not found"})
			return
		}

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to fetch author"})
		return
	}

	render.JSON(w, r, author)
}

func (h *AuthorHandler) ListAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.authorService.ListAuthors(r.Context())
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to list authors"})
		return
	}

	result := make([]model.AuthorResponse, 0, len(authors))
	for _, author := range authors {
		result = append(result, author)
	}

	render.JSON(w, r, result)
}

func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, ok := parseAuthorID(w, r)
	if !ok {
		return
	}

	if _, err := h.authorService.GetAuthor(r.Context(), authorID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": "author not found"})
			return
		}

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to fetch author"})
		return
	}

	var payload model.AuthorRequest
	if err := decodeJSON(r, &payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	updatedAuthor, err := h.authorService.UpdateAuthor(r.Context(), authorID, payload)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to update author"})
		return
	}

	render.JSON(w, r, updatedAuthor)
}

func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, ok := parseAuthorID(w, r)
	if !ok {
		return
	}

	if _, err := h.authorService.GetAuthor(r.Context(), authorID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, map[string]string{"error": "author not found"})
			return
		}

		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to fetch author"})
		return
	}

	err := h.authorService.DeleteAuthor(r.Context(), authorID)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "failed to delete author"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseAuthorID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	authorID, err := strconv.ParseInt(chi.URLParam(r, "authorID"), 10, 64)
	if err != nil || authorID <= 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid author id"})
		return 0, false
	}

	return authorID, true
}

func decodeJSON(r *http.Request, payload *model.AuthorRequest) error {
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return errors.New("invalid request body")
	}

	payload.Name = strings.TrimSpace(payload.Name)
	if payload.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
