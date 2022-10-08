package database

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/jailtonjunior94/go-tests/internal/entities"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestClientRepository_Save(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

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
	assert.Nil(t, err)

	clientRepository := NewClientRepository(db)
	client := &entities.Client{
		ID:     "123",
		Name:   "Jailton Junior",
		Email:  "jailton.junior94@outlook.com",
		Points: 0,
	}

	err = clientRepository.Save(client)
	assert.Nil(t, err)

	var id, name, email string
	var points int

	err = db.QueryRow("SELECT id, name, email, points FROM clients WHERE id = $1", client.ID).Scan(&id, &name, &email, &points)
	assert.Nil(t, err)

	assert.Equal(t, "123", id)
}

func BenchmarkClientRepositoryMemory_Save(b *testing.B) {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(b, err)

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
	assert.Nil(b, err)

	clientRepository := NewClientRepository(db)

	for i := 0; i < b.N; i++ {
		client := &entities.Client{
			ID:     uuid.New().String(),
			Name:   "Jailton Junior",
			Email:  "jailton.junior94@outlook.com",
			Points: 0,
		}

		clientRepository.Save(client)
	}
}

func BenchmarkClientRepositoryDisk_Save(b *testing.B) {
	db, err := sql.Open("sqlite3", "test.db")
	assert.Nil(b, err)

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
	assert.Nil(b, err)

	clientRepository := NewClientRepository(db)

	for i := 0; i < b.N; i++ {
		client := &entities.Client{
			ID:     uuid.New().String(),
			Name:   "Jailton Junior",
			Email:  "jailton.junior94@outlook.com",
			Points: 0,
		}

		clientRepository.Save(client)
	}
}
