package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Class struct {
	mongox.Model
	TableName          string             `bson:"table"`
	Name               string             `bson:"name"`
	HitDie             int                `bson:"hit_die"`
	ProficiencyChoices ProficiencyChoices `bson:"proficiency_choices"`
	ClassLevels        ClassLevelsApiUrl  `bson:"class_levels"`
	Subclass           Subclasses         `bson:"subclass"`
}

type ProficiencyChoices struct{}

type Subclasses struct{}

type ClassLevelsApiUrl string

func (c *Class) GetTableName() string {
	return c.TableName
}

func (c *Class) NewClassCollection(db *mongo.Database) *mongox.Collection[Class] {
	return mongox.NewCollection[Class](db.Collection(c.GetTableName()))
}
