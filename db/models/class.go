package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Class struct {
	mongox.Model
	Name string `bson:"name"`
}

func NewClassCollection(db *mongo.Database) *mongox.Collection[Class] {
	return mongox.NewCollection[Class](db.Collection("classes"))
}
