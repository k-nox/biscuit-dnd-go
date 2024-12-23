package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LanguagesModel struct {
	LangaugesSchema LanguagesSchema
	TableName       string
}

const languagesTableName = "languages"

type LanguagesSchema struct {
	mongox.Model
	Index           string   `bson:"index"`
	Name            string   `bson:"name"`
	Type            string   `bson:"type"`
	Script          string   `bson:"script"`
	ApiURL          string   `bson:"url"`
	TypicalSpeakers []string `bson:"typical_speakers"`
}

func (l *LanguagesModel) NewLanguagesCollection(db *mongo.Database) *mongox.Collection[LanguagesSchema] {
	return mongox.NewCollection[LanguagesSchema](db.Collection(l.GetTableName()))
}

func (l *LanguagesModel) GetTableName() string {
	return languagesTableName
}

func (l *LanguagesModel) GetSchema() interface{} {
	return l.LangaugesSchema
}
