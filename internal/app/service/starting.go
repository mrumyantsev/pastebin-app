package service

import (
	"context"
	"time"

	"github.com/mrumyantsev/errlib-go"
	"github.com/rs/zerolog/log"
)

func (a *App) start(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.AppStartTimeoutSeconds)*time.Second)
	defer cancel()

	err := a.databases.Mongo.Connect(ctx)
	if err != nil {
		return errlib.Wrap(err, "could not connect to mongo database")
	}

	if err = a.databases.Postgres.Connect(ctx); err != nil {
		return errlib.Wrap(err, "could not connect to postgres database")
	}

	if a.config.DatabaseMigrate {
		if err = a.migrate(ctx); err != nil {
			return errlib.Wrap(err, "could not migrate")
		}
	}

	go func() {
		log.Info().
			Str("host", a.config.ServerHost).
			Int64("port", a.config.ServerPort).
			Msg("server started")

		if err = a.server.Run(); err != nil && !a.isShutdown {
			a.errorCh <- errlib.Wrap(err, "could not run server")
		}
	}()

	return nil
}
