package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SkillsModel struct {
	AbilityScoreSchema SkillsSchema
}

const SkillsTableName = "skills"

type SkillsSchema struct {
	mongox.Model
	Index        string       `bson:"index"`
	Name         string       `bson:"name"`
	Desc         []string     `bson:"desc"`
	AbilityScore AbilityScore `bson:"ability_score"`
}

type AbilityScore struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

func NewSkillsModel() *SkillsModel {
	return &SkillsModel{}
}

func (a *SkillsModel) NewSkillsCollection(db *mongo.Database) *mongox.Collection[SkillsSchema] {
	return mongox.NewCollection[SkillsSchema](db.Collection(a.GetTableName()))
}

func (a *SkillsModel) GetTableName() string {
	return SkillsTableName
}
