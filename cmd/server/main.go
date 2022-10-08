package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jailtonjunior94/go-tests/internal/infra/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	controller := controller.NewBaseHandler(db)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Post("/clients", controller.CreateClientHander)

	http.ListenAndServe(":8080", router)
}
