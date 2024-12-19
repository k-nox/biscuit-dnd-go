package models

import (
	"github.com/chenmingyong0423/go-mongox/v2"
	"github.com/k-nox/biscuit-dnd-go/db"
)

type Class struct {
	mongox.Model
	Name string `bson:"name"`
}

func NewClassCollection() *mongox.Collection[Class] {
	mongoColl := db.GetCollection("classes")
	return mongox.NewCollection[Class](mongoColl)
}
