package paste

import (
	"context"

	"github.com/mrumyantsev/go-base64conv"
	"github.com/mrumyantsev/pastebin-app/internal/database"
)

type DummyHttpAdapter struct {
	db *database.PostgresDatabase
}

func NewDummyHttpAdapter(db *database.PostgresDatabase) *DummyHttpAdapter {
	return &DummyHttpAdapter{
		db: db,
	}
}

func (a *DummyHttpAdapter) GetGeneratedPasteId(ctx context.Context) (string, error) {
	const query = "SELECT COALESCE(MAX(id), 0) + 1 FROM pastes"

	row := a.db.QueryRow(ctx, query)

	var id int64

	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return base64conv.ItobRawUrl(id), nil
}
