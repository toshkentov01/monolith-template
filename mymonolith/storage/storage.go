package storage

import (
	"gitlab.com/sardortoshkentov/mymonolith/storage/postgres"
	"gitlab.com/sardortoshkentov/mymonolith/storage/repo"
)

// StorageInterface is an interface for storage
type StorageInterface interface {
	User() repo.UserRepository
}

type storage struct {
	userRepo repo.UserRepository
}

// NewStorage ...
func NewStorage() StorageInterface {
	return &storage{
		userRepo: postgres.NewUserRepo(),
	}
}

func (s storage) User() repo.UserRepository {
	return s.userRepo
}
