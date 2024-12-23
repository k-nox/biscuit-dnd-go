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
	class := models.ClassModel{}
	classColl := class.NewClassCollection(a.db)
	classResult, err := db.FindOneByKey(context.Background(), classColl, "name", "Barbarian")

	if err != nil {
		panic(err)
	}

	fmt.Println("id : ", classResult.ID, " name : ", classResult.Name)

	lang := models.LanguagesModel{}
	langColl := lang.NewLanguagesCollection(a.db)
	langResult, err := db.FindOneByKey(context.Background(), langColl, "name", "Primordial")

	if err != nil {
		panic(err)
	}

	fmt.Println("id : ", langResult.ID, " name : ", langResult.Name, " typical speakers : ", langResult.TypicalSpeakers)
}
