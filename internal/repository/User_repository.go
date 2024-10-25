package repository

import (
	"errors"
	"sync"

	"github.com/yordil/mereb-tech-challenge/internal/domain"
)

type UserRepository struct {
	userrepo map[string]domain.User
	mu       sync.RWMutex
}


func NewUserRepository() domain.UserRepository {
	return &UserRepository{
		userrepo: make(map[string]domain.User),
	}
}

// CreateUser saves a user to the repository write operation
func (pr *UserRepository) CreateUser(user domain.User) error {
	pr.mu.Lock()         
	defer pr.mu.Unlock() 
	pr.userrepo[user.ID] = user
	return nil
}

// GetAll retrieves all users from the repository read operation
func (pr *UserRepository) GetAll() ([]domain.User, error) {
	pr.mu.RLock()        
	defer pr.mu.RUnlock()

	var users []domain.User
	for _, user := range pr.userrepo {
		users = append(users, user)
	}
	return users, nil
}


func (pr *UserRepository) GetUserById(id string) (domain.User, error) {
	pr.mu.RLock()         
	defer pr.mu.RUnlock() 

	user, ok := pr.userrepo[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

// DeleteUser removes a user from the repository by ID write operation.
func (pr *UserRepository) DeleteUser(id string) error {
	pr.mu.Lock()         
	defer pr.mu.Unlock() 
	delete(pr.userrepo, id)
	return nil
}

// UpdateUser updates a user's information in the repository (write operation).
func (pr *UserRepository) UpdateUser(id string, user domain.User) (domain.User, error) {
	pr.mu.Lock()         
	defer pr.mu.Unlock() 
	pr.userrepo[id] = user
	return user, nil
}
