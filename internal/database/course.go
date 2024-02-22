package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryID string) (*Course, error){
	id := uuid.New().String()
	_, err := (`INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)`, id, name, description, categoryID)
	err := c.db.Exec(query,id, name, description, categoryID).Scan(&c.ID)
	if err != nil {
		return nil, err
	}

	return &Course{	
		ID: id,
		Name: name,
		Description: description,
		CategoryID: categoryID,
	}, nil
	
}