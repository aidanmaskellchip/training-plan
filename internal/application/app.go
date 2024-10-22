package application

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"training-plan/internal/infrastructure/config"
	"training-plan/internal/infrastructure/database/db"
	"training-plan/internal/infrastructure/event"
	"training-plan/internal/infrastructure/repository"
)

type App struct {
	Config        *config.Config
	Logger        *log.Logger
	DB            *gorm.DB
	Repos         *repository.Repositories
	EventBus      *event.EventBus
	EventChannels *event.Channels
}

func Load(
	conf *config.Config,
	logger *log.Logger,
	db *gorm.DB,
	repos *repository.Repositories,
	eb *event.EventBus,
	ec *event.Channels,
) (app App) {
	app.Config = conf
	app.Logger = logger
	app.DB = db
	app.Repos = repos
	app.EventBus = eb
	app.EventChannels = ec

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

	eBus := event.NewEventBus()
	eBus.LoadSubscriptions()

	return Load(
		conf,
		logger,
		database,
		repository.NewRepos(database),
		eBus,
		event.NewChannels(),
	), nil
}
