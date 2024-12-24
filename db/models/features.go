package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FeaturesModel struct {
	FeaturesSchema FeaturesSchema
}

const featureTableName = "features"

type FeaturesSchema struct {
	mongox.Model
	Index  string   `bson:"index"`
	Name   string   `bson:"name"`
	Class  Class    `bson:"class"`
	Level  int      `bson:"level"`
	Desc   []string `bson:"desc"`
	APIURL string   `bson:"url"`
}

type Class struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

func NewFeatureModel() *EquipmentModel {
	return &EquipmentModel{}
}

func (f *FeaturesModel) NewFeaturesCollection(db *mongo.Database) *mongox.Collection[EquipmentSchema] {
	return mongox.NewCollection[EquipmentSchema](db.Collection(f.GetTableName()))
}

func (f *FeaturesModel) GetTableName() string {
	return featureTableName
}

func (f *FeaturesModel) GetSchema() interface{} {
	return f.FeaturesSchema
}
