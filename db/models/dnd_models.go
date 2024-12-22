package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// All reusable common methods should be defined here
// All dnd models should implement this interface

type DndModels interface {
	NewClassCollection(db *mongo.Database) *mongox.Collection[Class]
	GetTableName() string
}
