package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	exchangemodels "gitlab.com/sardortoshkentov/mymonolith/exchange_models"
	"gitlab.com/sardortoshkentov/mymonolith/pkg/sqls"
	"gitlab.com/sardortoshkentov/mymonolith/platform/postgres"
	"gitlab.com/sardortoshkentov/mymonolith/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo() repo.UserRepository {

	return &userRepo{
		db: postgres.DB(),
	}
}

// CreateUser methods creates user
func (ur *userRepo) CreateUser(user *exchangemodels.CreateUserModel) error {
	if ur.db == nil{
		log.Println("DB is NILL")
	}
	_, err := ur.db.Exec(sqls.InsertUser, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
