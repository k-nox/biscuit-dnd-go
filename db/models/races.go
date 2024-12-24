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
	Index               string                `bson:"index"`
	Name                string                `bson:"name"`
	Desc                string                `bson:"desc"`
	Speed               int                   `bson:"speed"`
	Age                 string                `bson:"age"`
	Alignment           string                `bson:"alignment"`
	SizeDescription     string                `bson:"size_description"`
	Size                string                `bson:"size"`
	StartingProficiency []StartingProficiency `bson:"starting_proficiencies"`
	Languages           []Languages           `bson:"languages"`
	LanguageDescription string                `bson:"language_desc"`
	Traits              []Traits              `bson:"traits"`
	Subraces            []Subraces            `bson:"subraces"`
	APIURL              string                `bson:"url"`
}

type StartingProficiency struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

type Languages struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

type Traits struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

type Subraces struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
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
