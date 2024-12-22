package config

import (
	"github.com/k-nox/biscuit-dnd-go/db"
)

type Config struct {
	Database db.Config `yaml:"database"`
}
