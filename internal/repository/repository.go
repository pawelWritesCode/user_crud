package repository

import "example-user-crud/internal/models"

type UsersRepository interface {
	Create(u models.User) (models.User, error)
	Delete(userId int) error
	Replace(userId int, to models.User) error
	FindOne(userId int) (models.User, error)
	GetAll() []models.User
}
