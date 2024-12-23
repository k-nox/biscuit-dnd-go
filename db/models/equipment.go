package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EquipmentModel struct {
	LangaugesSchema EquipmentSchema
}

const tableName = "languages"

type EquipmentSchema struct {
	mongox.Model
	Index           string   `bson:"index"`
	Name            string   `bson:"name"`
	Type            string   `bson:"type"`
	Script          string   `bson:"script"`
	url             string   `bson:"url"`
	TypicalSpeakers []string `bson:"typical_speakers"`
}

func (l *EquipmentModel) NewEquipmentCollection(db *mongo.Database) *mongox.Collection[EquipmentSchema] {
	return mongox.NewCollection[EquipmentSchema](db.Collection(l.GetTableName()))
}

func (l *EquipmentModel) GetTableName() string {
	return tableName
}

func (l *EquipmentModel) GetSchema() interface{} {
	return l.LangaugesSchema
}
