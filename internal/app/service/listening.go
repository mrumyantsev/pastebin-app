package service

import "github.com/rs/zerolog/log"

func (a *App) awaitSignalOrError() error {
	select {
	case sig := <-a.signalCh:
		a.isShutdown = true

		log.Info().
			Str("signal", sig.String()).
			Msg("stop signal received. shutting down...")
	case err := <-a.errorCh:
		return err
	}

	return nil
}
