package app

import (
	"context"
	"fmt"
	"io"

	"github.com/k-nox/biscuit-dnd-go/config"
	"github.com/k-nox/biscuit-dnd-go/db"
	"github.com/k-nox/biscuit-dnd-go/db/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	db     *mongo.Database
	writer io.Writer
}

func New(c config.Config, w io.Writer) (*App, error) {
	db, err := db.New(c.Database)
	if err != nil {
		return nil, fmt.Errorf("error with db: %w", err)
	}

	return &App{
		db:     db,
		writer: w,
	}, nil
}

func (a *App) Run() error {
	classColl := models.NewClassCollection(a.db)
	class, err := db.FindOneByKey(context.Background(), classColl, "name", "Barbarian")

	if err != nil {
		return fmt.Errorf("error performing lookup: %w", err)
	}

	fmt.Fprintln(a.writer, class.Name)
	return nil
}
