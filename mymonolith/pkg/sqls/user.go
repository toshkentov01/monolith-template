package sqls

const (
	// InsertUser ...
	InsertUser = `
		INSERT INTO users (id, username, email, password) VALUES($1, $2, $3, crypt($4, gen_salt('bf')))
	`
)