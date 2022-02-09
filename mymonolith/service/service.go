package service

import (
	"context"
	"errors"

	"gitlab.com/sardortoshkentov/mymonolith/config"
	exchangemodels "gitlab.com/sardortoshkentov/mymonolith/exchange_models"
	l "gitlab.com/sardortoshkentov/mymonolith/pkg/logger"
	"gitlab.com/sardortoshkentov/mymonolith/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pkgLogger "gitlab.com/sardortoshkentov/mymonolith/pkg/logger"
)

var (
	//ErrEmailIncorrect ...
	ErrEmailIncorrect = errors.New("email is incorrect")
)

// UserService ...
type UserService struct {
	storage storage.StorageInterface
	logger  l.Logger
}

// NewUserService ...
func NewUserService(log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStorage(),
		logger:  log,
	}
}

// User ...
func User() *UserService {
	cfg := config.Config()

	logPkg := pkgLogger.New(cfg.LogLevel, "mymonolith")
	defer pkgLogger.Cleanup(logPkg)

	return &UserService{
		storage: storage.NewStorage(),
		logger: logPkg,
	}
}

// CreateUser creates a user
func (ur *UserService) CreateUser(ctx context.Context, req *exchangemodels.CreateUserModel) (*exchangemodels.EmptyModel, error) {
	err := ur.storage.User().CreateUser(req)
	if err != nil {
		ur.logger.Error("Error while creating a user, ERROR: " + err.Error())
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	return &exchangemodels.EmptyModel{}, nil
}
