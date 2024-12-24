package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClassModel struct {
	ClassSchema ClassSchema
}

const classTableName = "classes"

// TODO: split up work on class into multiple different tasks, this file will prolly get super big

type ClassSchema struct {
	mongox.Model
	Name              string              `bson:"name"`
	ClassLevels       string              `bson:"class_levels"`
	Subclasses        []Subclass          `bson:"subclasses"`
	Proficiencies     []Proficiencies     `bson:"proficiencies"`
	SavingThrows      []SavingThrows      `bson:"saving_throws"`
	StartingEquipment []StartingEquipment `bson:"starting_equipment"`
}

type Subclass struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

type Proficiencies struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

type SavingThrows struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

type StartingEquipment struct {
	Equipment Equipment
}

type Equipment struct {
	Index  string `bson:"index"`
	Name   string `bson:"name"`
	APIURL string `bson:"url"`
}

func NewClassModel() *ClassModel {
	return &ClassModel{}
}

func (c *ClassModel) NewClassCollection(db *mongo.Database) *mongox.Collection[ClassSchema] {
	return mongox.NewCollection[ClassSchema](db.Collection(c.GetTableName()))
}

func (c *ClassModel) GetTableName() string {
	return classTableName
}

func (c *ClassModel) GetSchema() interface{} {
	return c.ClassSchema
}
