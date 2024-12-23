package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EquipmentModel struct {
	EquipmentSchema EquipmentSchema
}

const equipmentTableName = "equipment"

type EquipmentSchema struct {
	mongox.Model
	Index             string            `bson:"index"`
	Name              string            `bson:"name"`
	WeaponCategory    string            `bson:"weapon_category"`
	WeaponRange       string            `bson:"weapon_range"`
	CategoryRange     string            `bson:"category_range"`
	Weight            int               `bson:"weight"`
	Url               string            `bson:"url"`
	EquipmentCategory EquipmentCategory `bson:"equipment_category"`
}

type EquipmentCategory struct {
	Index string `bson:"index"`
	Name  string `bson:"name"`
	Url   string `bson:"url"`
}

func (l *EquipmentModel) NewEquipmentCollection(db *mongo.Database) *mongox.Collection[EquipmentSchema] {
	return mongox.NewCollection[EquipmentSchema](db.Collection(l.GetTableName()))
}

func (l *EquipmentModel) GetTableName() string {
	return equipmentTableName
}

func (l *EquipmentModel) GetSchema() interface{} {
	return l.EquipmentSchema
}
