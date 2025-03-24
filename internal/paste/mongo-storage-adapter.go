package paste

import (
	"context"
	"errors"

	"github.com/mrumyantsev/pastebin-app/internal/database"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	optionsUpsert = options.UpdateOne().SetUpsert(true)
)

type mongoPaste struct {
	id      int64  `bson:"_id"`
	content []byte `bson:"content"`
}

type MongoStorageAdapter struct {
	db *database.MongoDatabase
}

func NewMongoStorageAdapter(db *database.MongoDatabase) *MongoStorageAdapter {
	return &MongoStorageAdapter{
		db: db,
	}
}

func (a *MongoStorageAdapter) CreatePasteContentById(ctx context.Context, id int64, content []byte) error {
	doc := mongoPaste{
		id:      id,
		content: content,
	}

	_, err := a.db.Collection("pastes").InsertOne(ctx, doc)
	if err != nil {
		return err
	}

	return nil
}

func (a *MongoStorageAdapter) CreateOrUpdatePasteContentById(ctx context.Context, id int64, content []byte) error {
	update := mongoPaste{
		id:      id,
		content: content,
	}

	_, err := a.db.Collection("pastes").UpdateByID(ctx, id, update, optionsUpsert)
	if err != nil {
		return err
	}

	return nil
}

func (a *MongoStorageAdapter) GetPasteContentById(ctx context.Context, id int64) ([]byte, error) {
	filter := mongoPaste{
		id: id,
	}

	res := a.db.Collection("pastes").FindOne(ctx, filter)

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrPasteNotFound
	}
	if err != nil {
		return nil, err
	}

	var paste mongoPaste

	if err = res.Decode(&paste); err != nil {
		return nil, err
	}

	return paste.content, nil
}

func (a *MongoStorageAdapter) DeletePasteContentById(ctx context.Context, id int64) error {
	filter := mongoPaste{
		id: id,
	}

	_, err := a.db.Collection("pastes").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (a *MongoStorageAdapter) IsPasteContentExistsById(ctx context.Context, id int64) (bool, error) {
	filter := mongoPaste{
		id: id,
	}

	res := a.db.Collection("pastes").FindOne(ctx, filter)

	err := res.Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
