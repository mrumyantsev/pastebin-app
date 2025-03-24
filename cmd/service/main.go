package main

import (
	"context"

	"github.com/mrumyantsev/pastebin-app/internal/app/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
}

func main() {
	ctx := context.Background()

	app, err := service.NewApp(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize service")
	}

	if err = app.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to run service")
	}
}
