package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Languages struct {
	ClassSchema ClassSchema
}

const tableName = "languages"

type LanguagesSchema struct {
	mongox.Model
	Index string `bson:"index"`
	Name  string `bson:"name"`
	Type  string `bson:"type"`
}

func (c *Languages) NewClassCollection(db *mongo.Database) *mongox.Collection[ClassSchema] {
	return mongox.NewCollection[ClassSchema](db.Collection(c.GetTableName()))
}

func (c *Languages) GetTableName() string {
	return tableName
}
