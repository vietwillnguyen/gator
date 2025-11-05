package commands

import (
	"fmt"
	"gator/internal/config"
)

type State struct {
	config *config.Config
}

// Command represents a CLI command
type Command struct {
	Name        string
	Arguments   []string
	Description string
	Callback    func(*State, Command) error
}

var CommandsMap map[string]Command

func CommandHelp(config *State, cmd Command) error {
	fmt.Println("Usage:")
	// Find longest command name for padding
	maxLen := 0
	for _, cmd := range CommandsMap {
		if len(cmd.Name) > maxLen {
			maxLen = len(cmd.Name)
		}
	}
	// Print aligned
	for _, cmd := range CommandsMap {
		fmt.Printf("  %-*s  %s\n", maxLen, cmd.Name, cmd.Description)
	}

	return nil
}

func CommandLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("CommandLogin: expects a single argument, the username.")
	}
	err := s.config.SetUser(cmd.Arguments[0])
	if err != nil {
		return err
	}
	fmt.Println("CommandLogin: User has been set.")
	return nil
}
