package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jailtonjunior94/go-tests/internal/infra/controller"
	_ "github.com/mattn/go-sqlite3"
)

func setupDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
		CREATE TABLE IF NOT EXISTS clients 
		(
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			email TEXT,
			points INTEGER
		);
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	return db
}

func tearDownDb(db *sql.DB) {
	db.Exec("DROP TABLE clients")
	db.Close()
}

func TestCreateClientHandler(t *testing.T) {
	db := setupDb()
	defer tearDownDb(db)

	controller := controller.NewBaseHandler(db)
	t.Run("should create a client", func(t *testing.T) {
		data := `{ "name":"Jailton Junior", "email":"jailton.junior94@outlook.com" }`
		reader := strings.NewReader(data)

		request, _ := http.NewRequest("POST", "/clients", reader)
		response := httptest.NewRecorder()

		controller.CreateClientHander(response, request)
		if response.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, response.Code)
		}
	})
}
