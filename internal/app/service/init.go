package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/mrumyantsev/errlib-go"
	"github.com/mrumyantsev/pastebin-app/internal/auth"
	"github.com/mrumyantsev/pastebin-app/internal/database"
	"github.com/mrumyantsev/pastebin-app/internal/paste"
	"github.com/mrumyantsev/pastebin-app/internal/server"
	"github.com/mrumyantsev/pastebin-app/internal/user"
)

func (a *App) init(_ context.Context) error {
	// Config
	a.config = NewConfig()

	err := a.config.Init()
	if err != nil {
		return errlib.Wrap(err, "could not initialize application config")
	}

	// Logger
	if err = a.initLogger(); err != nil {
		return errlib.Wrap(err, "could not initialize logger")
	}

	// Databases
	a.databases.Postgres, err = database.NewPostgresDatabase(&database.PostgresConfig{
		Host:                   a.config.PostgresDatabaseHost,
		Port:                   a.config.PostgresDatabasePort,
		Username:               a.config.PostgresDatabaseUsername,
		Password:               a.config.PostgresDatabasePassword,
		Name:                   a.config.PostgresDatabaseName,
		MaxConns:               a.config.PostgresDatabaseMaxConns,
		MaxConnIdleTimeMinutes: a.config.PostgresDatabaseMaxConnIdleTimeMinutes,
	})
	if err != nil {
		return errlib.Wrap(err, "could not initialize postgres database config")
	}

	a.databases.Mongo, err = database.NewMongoDatabase(&database.MongoConfig{
		Host:     a.config.MongoDatabaseHost,
		Port:     a.config.MongoDatabasePort,
		Username: a.config.MongoDatabaseUsername,
		Password: a.config.MongoDatabasePassword,
		Name:     a.config.MongoDatabaseName,
	})
	if err != nil {
		return errlib.Wrap(err, "could not initialize mongo database config")
	}

	// Domain Adapters
	userDatabaseAdapter := user.NewPostgresDatabaseAdapter(a.databases.Postgres)
	pasteDatabaseAdapter := paste.NewPostgresDatabaseAdapter(a.databases.Postgres)
	pasteStorageAdapter := paste.NewMongoStorageAdapter(a.databases.Mongo)
	pasteHttpAdapter := paste.NewDummyHttpAdapter(a.databases.Postgres)

	// Domain Services
	a.services.Auth = auth.NewService(userDatabaseAdapter)
	a.services.User = user.NewService(userDatabaseAdapter)
	a.services.Paste = paste.NewService(pasteDatabaseAdapter, pasteStorageAdapter, pasteHttpAdapter)

	// Server
	a.server = server.NewHttpServer(&server.HttpConfig{
		Host:                a.config.ServerHost,
		Port:                a.config.ServerPort,
		ReadTimeoutSeconds:  a.config.ServerReadTimeoutSeconds,
		WriteTimeoutSeconds: a.config.ServerWriteTimeoutSeconds,
		MaxHeaderBytes:      a.config.ServerMaxHeaderBytes,
		IsEnableDebugMode:   a.config.ServerEnableDebugMode,
	})

	// Routes v1
	if err = a.initRoutesV1(); err != nil {
		return errlib.Wrap(err, "could not initialize v1 routes")
	}

	// Shutdown
	a.errorCh = make(chan error)
	a.signalCh = make(chan os.Signal, 1)

	signal.Notify(a.signalCh, syscall.SIGINT, syscall.SIGTERM)

	a.isShutdown = false

	return nil
}
