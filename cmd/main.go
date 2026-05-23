package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	ah "github.com/eugbyte/sqlc_tutorial/internal/api/author/handler"
	authrepo "github.com/eugbyte/sqlc_tutorial/internal/api/author/repository/codegen"
	"github.com/eugbyte/sqlc_tutorial/internal/api/author/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	// "postgres://username:password@localhost:5432/database_name"
	const username = "postgres"
	const password = "postgres"
	const host = "localhost"
	const dbName = "sqlc_tutorial"
	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, 5432, dbName))
	if err != nil {
		log.Fatalf("pgx.Connect(): %v", err)
	}
	defer conn.Close(ctx)

	authorRepo := authrepo.New(conn)
	authorService := service.NewAuthorService(authorRepo)
	authorHandler := ah.NewAuthorHandler(authorService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"message": "Hello, World!"}
		render.JSON(w, r, data)
	})
	ah.RegisterRoutes(router, authorHandler)

	log.Println("server starting at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
