package paste

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mrumyantsev/pastebin-app/internal/database"
)

type PostgresDatabaseAdapter struct {
	db *database.PostgresDatabase
}

func NewPostgresDatabaseAdapter(db *database.PostgresDatabase) *PostgresDatabaseAdapter {
	return &PostgresDatabaseAdapter{
		db: db,
	}
}

func (a *PostgresDatabaseAdapter) CreatePaste(ctx context.Context, paste Paste) error {
	const query = "INSERT INTO pastes (id, user_id, content, expires_at) VALUES ($1, $2, $3, $4)"

	_, err := a.db.Exec(ctx, query, paste.Id, paste.UserId, paste.Content, paste.ExpiresAt)
	if err != nil {
		return err
	}

	return nil
}

func (a *PostgresDatabaseAdapter) GetAllPastes(ctx context.Context) ([]Paste, error) {
	const query = "SELECT id, user_id, created_at, expires_at FROM pastes"

	rows, err := a.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	pastes := []Paste{}
	var paste Paste

	for rows.Next() {
		if err = rows.Scan(&paste.Id, &paste.UserId, &paste.CreatedAt, &paste.ExpiresAt); err != nil {
			return nil, err
		}

		pastes = append(pastes, paste)
	}

	return pastes, nil
}

func (a *PostgresDatabaseAdapter) GetPasteById(ctx context.Context, id int64) (Paste, error) {
	const query = "SELECT id, user_id, content, created_at, expires_at FROM pastes WHERE id = $1"

	row := a.db.QueryRow(ctx, query, id)

	var paste Paste

	err := row.Scan(&paste.Id, &paste.UserId, &paste.Content, &paste.CreatedAt, &paste.ExpiresAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Paste{}, ErrPasteNotFound
	}
	if err != nil {
		return Paste{}, err
	}

	return paste, nil
}

func (a *PostgresDatabaseAdapter) UpdatePasteById(ctx context.Context, id int64, paste Paste) error {
	const query = "UPDATE pastes SET content = $1, expires_at = $2 WHERE id = $3"

	_, err := a.db.Exec(ctx, query, paste.Content, paste.ExpiresAt, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *PostgresDatabaseAdapter) DeletePasteById(ctx context.Context, id int64) error {
	const query = "DELETE FROM pastes WHERE id = $1"

	_, err := a.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
