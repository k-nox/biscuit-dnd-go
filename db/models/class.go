package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClassModel struct {
	ClassSchema ClassSchema
}

const classTableName = "classes"

type ClassSchema struct {
	mongox.Model
	Name        string `bson:"name"`
	ClassLevels string `bson:"class_levels"`
}

func (c *ClassModel) NewClassCollection(db *mongo.Database) *mongox.Collection[ClassSchema] {
	return mongox.NewCollection[ClassSchema](db.Collection(c.GetTableName()))
}

func (c *ClassModel) GetTableName() string {
	return classTableName
}
