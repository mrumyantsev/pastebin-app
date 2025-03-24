package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoConfig struct {
	Host     string
	Port     int64
	Username string
	Password string
	Name     string
}

type MongoDatabase struct {
	*mongo.Database
	config *MongoConfig
	client *mongo.Client
}

func NewMongoDatabase(cfg *MongoConfig) (*MongoDatabase, error) {
	return &MongoDatabase{
		config: cfg,
	}, nil
}

func (d *MongoDatabase) Connect(_ context.Context) error {
	connUri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s?authSource=admin",
		d.config.Username,
		d.config.Password,
		d.config.Host,
		d.config.Port,
		d.config.Name,
	)

	var err error

	d.client, err = mongo.Connect(options.Client().ApplyURI(connUri))
	if err != nil {
		return err
	}

	d.Database = d.client.Database(d.config.Name)

	return nil
}

func (d *MongoDatabase) Disconnect(ctx context.Context) error {
	err := d.client.Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
