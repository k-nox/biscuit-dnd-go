package main

import (
	"context"
	"fmt"

	"github.com/chenmingyong0423/go-mongox/v2/builder/query"
	"github.com/k-nox/biscuit-dnd-go/db/models"
)

func main() {
	classColl := models.NewClassCollection()
	res, err := classColl.Finder().Filter(query.NewBuilder().Eq("name", "Barbarian").Build()).FindOne(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
