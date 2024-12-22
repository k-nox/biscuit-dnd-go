package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DndModels interface {
	NewClassCollection(db *mongo.Database) *mongox.Collection[Class]
}
