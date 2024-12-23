package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Languages struct {
	LangaugesSchema LanguagesSchema
}

const tableName = "languages"

type LanguagesSchema struct {
	mongox.Model
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	Type   string `bson:"type"`
	Script string `bson:"script"`
	url    string `bson:"url"`
}

func (l *Languages) NewLanguagesCollection(db *mongo.Database) *mongox.Collection[LanguagesSchema] {
	return mongox.NewCollection[LanguagesSchema](db.Collection(l.GetTableName()))
}

func (l *Languages) GetTableName() string {
	return tableName
}

func (l *Languages) GetSchema() interface{} {
	return l.LangaugesSchema
}
