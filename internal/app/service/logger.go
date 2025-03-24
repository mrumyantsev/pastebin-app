package service

import (
	"strings"

	"github.com/rs/zerolog"
)

func (a *App) initLogger() error {
	var zerologLevel zerolog.Level

	switch strings.ToLower(a.config.LoggerGlobalLevel) {
	case "trace":
		zerologLevel = zerolog.TraceLevel
	case "debug":
		zerologLevel = zerolog.DebugLevel
	case "warn":
		zerologLevel = zerolog.WarnLevel
	case "error":
		zerologLevel = zerolog.ErrorLevel
	case "fatal":
		zerologLevel = zerolog.FatalLevel
	case "panic":
		zerologLevel = zerolog.PanicLevel
	default:
		zerologLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(zerologLevel)

	return nil
}
