package models

// All reusable common methods should be defined here
// All dnd models should implement this interface

type DndModels interface {
	GetTableName() string
	GetSchema() interface{}
}

// Add core methods here
