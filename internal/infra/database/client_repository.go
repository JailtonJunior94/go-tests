package database

import (
	"database/sql"

	"github.com/jailtonjunior94/go-tests/internal/entities"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (c *ClientRepository) Save(client *entities.Client) error {
	_, err := c.db.Exec("INSERT INTO clients (id, name, email, points) VALUES ($1, $2, $3, $4)", client.ID, client.Name, client.Email, client.Points)
	if err != nil {
		return err
	}
	return nil
}
