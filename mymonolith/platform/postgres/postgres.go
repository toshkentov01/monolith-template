package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //driver for postgres
	"gitlab.com/sardortoshkentov/mymonolith/config"
)

//DB get database instance
func DB() *sqlx.DB {
	var db *sqlx.DB
	dsn := postgresDSN()
	db = sqlx.MustConnect("postgres", dsn)
	return db
}

func postgresDSN() string {
	// URL for PostgreSQL connection.
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config().PostgresHost,
		config.Config().PostgresPort,
		config.Config().PostgresUser,
		config.Config().PostgresPassword,
		config.Config().PostgresDatabase,
	)
}
