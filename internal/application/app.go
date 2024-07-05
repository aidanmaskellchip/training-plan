package application

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"training-plan/internal/config"
	"training-plan/internal/data/repository"
	"training-plan/internal/database/db"
)

type App struct {
	Config *config.Config
	Logger *log.Logger
	DB     *gorm.DB
	Repos  *repository.Repositories
}

func Load(conf *config.Config, logger *log.Logger, db *gorm.DB, repos *repository.Repositories) (app App) {
	app.Config = conf
	app.Logger = logger
	app.DB = db
	app.Repos = repos

	return
}

func Setup() (app App, err error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	conf, err := config.Get()
	if err != nil {
		return app, fmt.Errorf("unable to setup config: %s", err.Error())
	}

	database, err := db.ConnectDb(*conf)
	if err != nil {
		return app, fmt.Errorf("unable to connect to db: %s", err.Error())
	}

	return Load(
		conf,
		logger,
		database,
		repository.NewRepos(database),
	), nil
}
