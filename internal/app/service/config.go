package service

import "github.com/kelseyhightower/envconfig"

type Config struct {
	// App
	AppStartTimeoutSeconds int64 `envconfig:"APP_START_TIMEOUT_SECONDS" default:"5"`
	AppStopTimeoutSeconds  int64 `envconfig:"APP_STOP_TIMEOUT_SECONDS" default:"5"`

	// Logger
	LoggerGlobalLevel string `envconfig:"LOGGER_GLOBAL_LEVEL" default:"info"`

	// Database
	DatabaseMigrate bool `envconfig:"DATABASE_MIGRATE" default:"true"`

	// Postgres Database
	PostgresDatabaseHost                   string `envconfig:"POSTGRES_DATABASE_HOST" default:"localhost"`
	PostgresDatabasePort                   int64  `envconfig:"POSTGRES_DATABASE_PORT" default:"8000"`
	PostgresDatabaseUsername               string `envconfig:"POSTGRES_DATABASE_USERNAME" default:"postgres"`
	PostgresDatabasePassword               string `envconfig:"POSTGRES_DATABASE_PASSWORD"`
	PostgresDatabaseName                   string `envconfig:"POSTGRES_DATABASE_NAME" default:"pastebin"`
	PostgresDatabaseMaxConns               int64  `envconfig:"POSTGRES_DATABASE_MAX_CONNS" default:"10"`
	PostgresDatabaseMaxConnIdleTimeMinutes int64  `envconfig:"POSTGRES_DATABASE_MAX_CONN_IDLE_TIME_MINUTES" default:"30"`

	// Mongo Database
	MongoDatabaseHost     string `envconfig:"MONGO_DATABASE_HOST" default:"localhost"`
	MongoDatabasePort     int64  `envconfig:"MONGO_DATABASE_PORT" default:"8002"`
	MongoDatabaseUsername string `envconfig:"MONGO_DATABASE_USERNAME" default:"root"`
	MongoDatabasePassword string `envconfig:"MONGO_DATABASE_PASSWORD"`
	MongoDatabaseName     string `envconfig:"MONGO_DATABASE_NAME" default:"pastebin"`

	// Server
	ServerHost                string `envconfig:"SERVER_HOST" default:"localhost"`
	ServerPort                int64  `envconfig:"SERVER_PORT" default:"8010"`
	ServerReadTimeoutSeconds  int64  `envconfig:"SERVER_READ_TIMEOUT_SECONDS" default:"5"`
	ServerWriteTimeoutSeconds int64  `envconfig:"SERVER_WRITE_TIMEOUT_SECONDS" default:"5"`
	ServerMaxHeaderBytes      int64  `envconfig:"SERVER_MAX_HEADER_BYTES" default:"1048576"`
	ServerEnableDebugMode     bool   `envconfig:"SERVER_ENABLE_DEBUG_MODE" default:"false"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}

	return nil
}
