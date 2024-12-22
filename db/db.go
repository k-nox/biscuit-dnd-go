package db

import (
	"context"
	"fmt"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func New(c Config) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", c.Host, c.Port)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo: %w", err)
	}
	return client.Database(c.Database), nil
}

func FindOneByKey[T any](ctx context.Context, collection *mongox.Collection[T], key string, value any) (*T, error) {
	return collection.Finder().Filter(query.NewBuilder().Eq(key, value).Build()).FindOne(ctx)
}
