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
	return &AbilityScoreModel{}
}

func (a *AbilityScoreModel) NewAbilityScoreCollection(db *mongo.Database) *mongox.Collection[AbilityScoreSchema] {
	return mongox.NewCollection[AbilityScoreSchema](db.Collection(a.GetTableName()))
}

func (a *AbilityScoreModel) GetTableName() string {
	return AbilityScoreTableName
}
