package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AbilityScoreModel struct {
	AbilityScoreSchema AbilityScoreSchema
}

const AbilityScoreTableName = "ability-scores"

type AbilityScoreSchema struct {
	mongox.Model
	Index    string         `bson:"index"`
	Name     string         `bson:"name"`
	FullName string         `bson:"full_name"`
	Desc     []string       `bson:"desc"`
	Skills   []AbilityScore `bson:"skills"`
	APIURL   string         `bson:"Url"`
}

func NewAbilityScoreModel() *AbilityScoreModel {
	//lint:ignore U1000 Returning an empty struct pointer
	return &AbilityScoreModel{AbilityScoreSchema: AbilityScoreSchema{
		Model:    mongox.Model{},
		Index:    "",
		Name:     "",
		FullName: "",
		Desc:     nil,
		Skills:   nil,
		APIURL:   "",
	}}
}

func (a *AbilityScoreModel) NewAbilityScoreCollection(db *mongo.Database) *mongox.Collection[AbilityScoreSchema] {
	return mongox.NewCollection[AbilityScoreSchema](db.Collection(a.GetTableName()))
}

func (a *AbilityScoreModel) GetTableName() string {
	return AbilityScoreTableName
}

func (a *AbilityScoreModel) GetSchema() interface{} {
	return a.AbilityScoreSchema
}
