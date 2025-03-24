package service

import (
	"context"
	"os"

	"github.com/mrumyantsev/pastebin-app/internal/auth"
	"github.com/mrumyantsev/pastebin-app/internal/database"
	"github.com/mrumyantsev/pastebin-app/internal/paste"
	"github.com/mrumyantsev/pastebin-app/internal/server"
	"github.com/mrumyantsev/pastebin-app/internal/user"
	"github.com/rs/zerolog/log"
)

type Databases struct {
	Postgres *database.PostgresDatabase
	Mongo    *database.MongoDatabase
}

type Services struct {
	Auth  auth.Servicer
	User  user.Servicer
	Paste paste.Servicer
}

type App struct {
	config     *Config
	databases  Databases
	services   Services
	server     *server.HttpServer
	errorCh    chan error
	signalCh   chan os.Signal
	isShutdown bool
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.init(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	log.Info().
		Int("pid", os.Getpid()).
		Msg("service started")

	err := a.start(ctx)
	if err != nil {
		return err
	}

	if err = a.awaitSignalOrError(); err != nil {
		return err
	}

	if err = a.shutdown(ctx); err != nil {
		return err
	}

	log.Info().Msg("service shut down")

	return nil
}
