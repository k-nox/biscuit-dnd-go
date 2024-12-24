package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RacesModel struct {
	RacesSchema RacesSchema
}

const racesTableName = "races"

type RacesSchema struct {
	Index string `bson:"index"`
	Name  string `bson:"name"`
}

func NewRacesModel() *RacesModel {
	return &RacesModel{}
}

func (r *RacesModel) NewRacesCollection(db *mongo.Database) *mongox.Collection[RacesSchema] {
	return mongox.NewCollection[RacesSchema](db.Collection(r.GetTableName()))
}

func (r *RacesModel) GetTableName() string {
	return racesTableName
}

func (r *RacesModel) GetSchema() interface{} {
	return r.RacesSchema
}
