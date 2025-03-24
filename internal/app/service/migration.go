package service

import (
	"context"
)

func (a *App) migrate(ctx context.Context) error {
	const query = `
		CREATE TABLE IF NOT EXISTS users (
			id         SERIAL8   NOT NULL PRIMARY KEY,
			username   VARCHAR   NOT NULL UNIQUE,
			password   VARCHAR   NOT NULL,
			email      VARCHAR   NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS pastes (
			id         INT8      NOT NULL PRIMARY KEY,
			user_id    INT8               REFERENCES users (id),
			content    VARCHAR   NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			expires_at TIMESTAMP NOT NULL
		);
`

	_, err := a.databases.Postgres.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
