package utils

import (
	"fmt"
	"gitlab.com/sardortoshkentov/mymonolith/config"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	config := config.Config()
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.PostgresHost,
			config.PostgresPort,
			config.PostgresUser,
			config.PostgresPassword,
			config.PostgresDatabase,
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			":%d",
			config.ServerPort,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
