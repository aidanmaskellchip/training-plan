package dbdsn

import (
	"fmt"
)

func DSNGenerator(env, host, port, user, password, dbName string) string {
	ssl := ""
	if env == "local" {
		ssl = "?sslmode=disable"
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s%s",
		user,
		password,
		host,
		port,
		dbName,
		ssl,
	)

	return dsn
}
