package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"training-plan/internal/api/infrastructure/config"
	"training-plan/internal/api/infrastructure/database/db/db_dsn"
)

func ConnectDb(conf config.Config) (*gorm.DB, error) {
	dbURL := dbdsn.DSNGenerator(
		conf.Env,
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbPassword,
		conf.DbName,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}

	return db, nil
}
