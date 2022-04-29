package database

import "example-user-crud/internal/models"

type DB struct {
	Index int
	Users []models.User
}

func New() *DB {
	return &DB{Users: []models.User{}, Index: 0}
}
