package service

import (
	"context"
	"time"

	"github.com/mrumyantsev/go-errlib"
)

func (a *App) shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(a.config.AppStopTimeoutSeconds)*time.Second)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		return errlib.Wrap(err, "could not shutdown server")
	}

	if err = a.databases.Postgres.Disconnect(ctx); err != nil {
		return errlib.Wrap(err, "could not disconnect from postgres database")
	}

	if err = a.databases.Mongo.Disconnect(ctx); err != nil {
		return errlib.Wrap(err, "could not disconnect from mongo database")
	}

	return nil
}
