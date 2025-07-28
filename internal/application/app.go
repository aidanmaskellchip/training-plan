package application

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-aws/sqs"
	"github.com/ThreeDotsLabs/watermill/message"
	"gorm.io/gorm"
	"log"
	"os"
	"training-plan/internal/infrastructure/accolade"
	"training-plan/internal/infrastructure/config"
	"training-plan/internal/infrastructure/database/db"
	"training-plan/internal/infrastructure/event/pubsub"
	"training-plan/internal/infrastructure/repository"
)

type App struct {
	Config      *config.Config
	Logger      *log.Logger
	DB          *gorm.DB
	Repos       *repository.Repositories
	SqsSub      *sqs.Subscriber
	SqsPub      *sqs.Publisher
	EventRouter *message.Router
	AccoladeService AccoladeService
}

func Load(
	conf *config.Config,
	logger *log.Logger,
	db *gorm.DB,
	repos *repository.Repositories,
	sub *sqs.Subscriber,
	pub *sqs.Publisher,
	eventRouter *message.Router,
	accoladeService AccoladeService,
) (app App) {
	app.Config = conf
	app.Logger = logger
	app.DB = db
	app.Repos = repos
	app.SqsSub = sub
	app.SqsPub = pub
	app.EventRouter = eventRouter
	app.AccoladeService = accoladeService

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

	wtmLogger := watermill.NewStdLogger(false, false)

	sub, err := pubsub.NewSQSSubscriber(&wtmLogger)
	if err != nil {
		panic(err)
	}
	pub, err := pubsub.NewSQSPublisher(&wtmLogger)
	if err != nil {
		panic(err)
	}
	eventRouter, err := pubsub.NewSQSEventRouter(&wtmLogger, sub)
	if err != nil {
		panic(err)
	}

	return Load(
		conf,
		logger,
		database,
		repository.NewRepos(database),
		sub,
		pub,
		eventRouter,
		&accolade.MockAccoladeService{},
	), nil
}
