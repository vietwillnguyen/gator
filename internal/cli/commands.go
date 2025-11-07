package cli

import "gator/internal/models"

type Commands struct {
	handlers map[string]func(*models.State, Command) error
}

func NewCommands() *Commands {
	return &Commands{
		handlers: make(map[string]func(*models.State, Command) error),
	}
}

func (c *Commands) Register(name string, f func(*models.State, Command) error) {
	c.handlers[name] = f
}

func (c *Commands) Run(s *models.State, cmd Command) error {
	return nil
}
