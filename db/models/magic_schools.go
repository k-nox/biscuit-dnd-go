package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MagicSchoolsModel struct {
	MagicSchoolsSchema MagicSchoolsSchema
	TableName          string
}

const MagicSchoolTableName = "magic-schools"

type MagicSchoolsSchema struct {
	mongox.Model
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	Desc   string `bson:"desc"`
	APIURL string `bson:"url"`
}

func NewMagicSchoolsModel() *MagicSchoolsModel {
	return &MagicSchoolsModel{}
}

func (m *MagicSchoolsModel) NewMagicSchoolsCollection(db *mongo.Database) *mongox.Collection[MagicSchoolsSchema] {
	return mongox.NewCollection[MagicSchoolsSchema](db.Collection(m.GetTableName()))
}

func (m *MagicSchoolsModel) GetTableName() string {
	return MagicSchoolTableName
}

func (m *MagicSchoolsModel) GetSchema() interface{} {
	return m.MagicSchoolsSchema
}
