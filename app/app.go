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
		return nil, err
	}

	return &App{
		db: db,
	}, nil
}

func (a *App) Run() {
	classColl := models.NewClassCollection(a.db)
	class, err := db.FindOneByKey(context.Background(), classColl, "name", "Barbarian")

	if err != nil {
		panic(err)
	}

	fmt.Println(class.ID, class.Name)
}
