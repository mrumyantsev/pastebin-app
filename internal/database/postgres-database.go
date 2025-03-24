package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Host                   string
	Port                   int64
	Username               string
	Password               string
	Name                   string
	MaxConns               int64
	MaxConnIdleTimeMinutes int64
}

type PostgresDatabase struct {
	*pgxpool.Pool
	config *pgxpool.Config
}

func NewPostgresDatabase(cfg *PostgresConfig) (*PostgresDatabase, error) {
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	pgxCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pgxCfg.MaxConns = int32(cfg.MaxConns)
	pgxCfg.MaxConnIdleTime = time.Duration(cfg.MaxConnIdleTimeMinutes) * time.Minute

	return &PostgresDatabase{
		config: pgxCfg,
	}, nil
}

func (d *PostgresDatabase) Connect(ctx context.Context) error {
	var err error

	d.Pool, err = pgxpool.NewWithConfig(ctx, d.config)
	if err != nil {
		return err
	}

	return nil
}

func (d *PostgresDatabase) Disconnect(_ context.Context) error {
	d.Pool.Close()

	return nil
}
