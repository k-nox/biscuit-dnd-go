package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClassModel struct {
	ClassSchema ClassSchema
}

const classTableName = "classes"

// TODO: split up work on class into multiple different tasks, this file will prolly get super big

type ClassSchema struct {
	mongox.Model
	Name        string     `bson:"name"`
	ClassLevels string     `bson:"class_levels"`
	Subclasses  []Subclass `bson:"subclasses"`
}

type Subclass struct {
	Index string `bson:"index"`
	Name  string `bson:"name"`
	Url   string `bson:"url"`
}

func (c *ClassModel) NewClassCollection(db *mongo.Database) *mongox.Collection[ClassSchema] {
	return mongox.NewCollection[ClassSchema](db.Collection(c.GetTableName()))
}

func (c *ClassModel) GetTableName() string {
	return classTableName
}

func (c *ClassModel) GetSchema() interface{} {
	return c.ClassSchema
}
