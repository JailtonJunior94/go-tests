package main

import (
	"database/sql"
	"net/http"

	"github.com/jailtonjunior94/go-tests/internal/infra/controller"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	controller := controller.NewBaseHandler(db)
	http.HandleFunc("/clients", controller.CreateClientHander)
	http.ListenAndServe(":8080", nil)
}
