package manager

import (
	"fmt"
	"sync"

	"example-user-crud/internal/database"
	"example-user-crud/internal/models"
)

var mu sync.Mutex

type UsersManager struct {
	db *database.DB
}

func NewUsersManager(db *database.DB) UsersManager {
	return UsersManager{db: db}
}

func (um UsersManager) Create(u models.User) (models.User, error) {
	for _, user := range um.db.Users {
		if user.FirstName == u.FirstName && user.LastName == u.LastName {
			return models.User{}, fmt.Errorf("user exists in database")
		}
	}

	mu.Lock()
	defer mu.Unlock()
	um.db.Index++
	u.Id = um.db.Index

	um.db.Users = append(um.db.Users, u)

	return u, nil
}

func (um UsersManager) Delete(userId int) error {
	newUsers := make([]models.User, 0, len(um.db.Users))
	mu.Lock()
	defer mu.Unlock()
	for i, user := range um.db.Users {
		if user.Id == userId {

			newUsers = append(um.db.Users[:i], um.db.Users[i+1:]...)
			um.db.Users = newUsers

			return nil
		}
	}

	return fmt.Errorf("could not find in database user of id %+v", userId)
}

func (um UsersManager) Replace(userId int, to models.User) error {
	newUsers := make([]models.User, 0, len(um.db.Users))
	mu.Lock()
	defer mu.Unlock()
	for i, user := range um.db.Users {
		if user.Id == userId {
			newUsers = append(um.db.Users[:i], um.db.Users[i+1:]...)

			to.Id = userId
			newUsers = append(newUsers, to)
			um.db.Users = newUsers

			return nil
		}
	}

	return fmt.Errorf("could not find in database user of id %+v", userId)
}

func (um UsersManager) FindOne(userId int) (models.User, error) {
	mu.Lock()
	defer mu.Unlock()
	for _, user := range um.db.Users {
		if user.Id == userId {
			return user, nil
		}
	}

	return models.User{}, fmt.Errorf("user of id %d does not exist in database", userId)
}

func (um UsersManager) GetAll() []models.User {
	mu.Lock()
	defer mu.Unlock()
	return um.db.Users
}
