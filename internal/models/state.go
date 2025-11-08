package models

import (
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
	// Later: DB connection, etc.
}
