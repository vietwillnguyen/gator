package models

import "gator/internal/config"

type State struct {
	Config *config.Config
	// Later: DB connection, etc.
}
