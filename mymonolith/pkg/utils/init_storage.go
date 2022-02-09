package utils

import (
	"log"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/jmoiron/sqlx"
	"gitlab.com/sardortoshkentov/mymonolith/config"
	"gitlab.com/sardortoshkentov/mymonolith/platform/postgres"
	p "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//InitStorage intializes and establishes database connect and return pointer to database and gorm adapter for casbin
func InitStorage(config *config.Configuration) (*sqlx.DB, *gormadapter.Adapter) {

	dsn, err := ConnectionURLBuilder("postgres")
	if err != nil {
		log.Fatal("Error building database URL")
	}

	db := postgres.DB()
	gormDB, err := gorm.Open(p.Open(dsn))
	if err != nil {
		log.Println("Could not connect to db with gorm")
		panic(err)
	}

	adapter, err := gormadapter.NewAdapterByDB(gormDB)
	if err != nil {
		log.Println("Could not create new adapter")
		panic(err)
	}

	return db, adapter
}
