package repo

import (
	"errors"

	exchangemodels "gitlab.com/sardortoshkentov/mymonolith/exchange_models"
)

var (
	// ErrAlreadyExists ...
	ErrAlreadyExists = errors.New("already exists")
	// ErrInvalidField ...
	ErrInvalidField = errors.New("incorrect field")
)

// UserRepository is an interface for client storage
type UserRepository interface {
	Reader
	Writer
}

// Reader ...
type Reader interface {

}

// Writer ...
type Writer interface {
	CreateUser(user *exchangemodels.CreateUserModel) error
}

