package app

import (
	"context"
	"fmt"
	"github.com/k-nox/biscuit-dnd-go/config"
	"github.com/k-nox/biscuit-dnd-go/db"
	"github.com/k-nox/biscuit-dnd-go/db/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	db *mongo.Database
}

func New(c config.Config) (*App, error) {
	db, err := db.New(c.Database)
	if err != nil {
		return nil, fmt.Errorf("error with db: %w", err)
	}

	return &App{
		db: db,
	}, nil
}

func (a *App) Run() {
	class := models.ClassModel{ClassSchema: models.ClassSchema{}}
	classColl := class.NewClassCollection(a.db)
	classResult, err := db.FindOneByKey(context.Background(), classColl, "name", "Barbarian")

	if err != nil {
		panic(err)
	}

	fmt.Println(
		"id : ", classResult.ID,
		" name : ", classResult.Name,
		" subclass : ", classResult.Subclasses,
		" subclass name : ", classResult.Subclasses[0].Name)

	lang := models.LanguagesModel{LangaugesSchema: models.LanguagesSchema{}}
	langColl := lang.NewLanguagesCollection(a.db)
	langResult, err := db.FindOneByKey(context.Background(), langColl, "name", "Primordial")

	if err != nil {
		panic(err)
	}

	fmt.Println("id : ", langResult.ID, " name : ", langResult.Name, " typical speakers : ", langResult.TypicalSpeakers)

	equipment := models.EquipmentModel{EquipmentSchema: models.EquipmentSchema{}}
	equipmentColl := equipment.NewEquipmentCollection(a.db)
	equipmentResult, err := db.FindOneByKey(context.Background(), equipmentColl, "name", "Club")
	if err != nil {
		panic(err)
	}

	fmt.Println(
		"id : ", equipmentResult.ID,
		" name : ", equipmentResult.Name,
		" weapon category : ", equipmentResult.WeaponCategory,
		" equipment category name : ", equipmentResult.EquipmentCategory.Name)

	abilityScore := models.AbilityScoreModel{AbilityScoreSchema: models.AbilityScoreSchema{}}
	abilityScoreColl := abilityScore.NewAbilityScoreCollection(a.db)
	abilityScoreResult, err := db.FindOneByKey(context.Background(), abilityScoreColl, "full_name", "Intelligence")
	if err != nil {
		panic(err)
	}

	fmt.Println(
		"id : ", abilityScoreResult.ID,
		" name : ", abilityScoreResult.Name,
		" full name : ", abilityScoreResult.FullName,
		" desc : ", abilityScoreResult.Desc)

	for _, skill := range abilityScoreResult.Skills {
		fmt.Println("skill : ", skill.Name)
	}
}
