package models

import "github.com/chenmingyong0423/go-mongox/v2"

type FeaturesModel struct {
	FeaturesSchema FeaturesSchema
}

type FeaturesSchema struct {
	mongox.Model
	Index string `bson:"index"`
	Name  string `bson:"name"`
}
