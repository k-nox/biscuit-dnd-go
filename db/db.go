package db

import (
	"context"
	"fmt"
	"net"

	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Config struct {
	Host     string `yaml:"host" env:"MONGODB_HOST"`
	Port     string `yaml:"port" env:"MONGODB_PORT"`
	Database string `yaml:"database" env:"MONGODB_DATABASE"`
}

func New(c Config) (*mongo.Database, error) {
	uri := "mongodb://" + net.JoinHostPort(c.Host, c.Port)
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("error connecting to mongo: %w", err)
	}

	return client.Database(c.Database), nil
}

func FindOneByKey[T any](ctx context.Context, collection *mongox.Collection[T], key string, value any) (*T, error) {
	doc, err := collection.Finder().Filter(query.NewBuilder().Eq(key, value).Build()).FindOne(ctx)
	if err != nil {
		return nil, fmt.Errorf("error finding one %T by key %s with value %v: %w", doc, key, value, err)
	}

	return doc, nil
}
