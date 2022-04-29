package main

import (
	"example-user-crud/internal/database"
	"example-user-crud/internal/manager"
)

type Container struct {
	UsersManager manager.UsersManager
}

func newContainer() *Container {
	db := database.New()
	um := manager.NewUsersManager(db)

	return &Container{UsersManager: um}
}
