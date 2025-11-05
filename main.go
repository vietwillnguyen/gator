package main

import (
	"bufio"
	"fmt"
	"gator/internal/config"
	"log"
	"os"
	"strings"
)

func init() {
	// CommandsMap holds all available commands
	CommandsMap = map[string]Command{
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    CommandHelp,
		},
		"login": {
			Name:        "login <username>",
			Description: "Log in with <username>",
			Callback:    CommandLogin,
		},
	}
}

// CleanInput normalizes user input by lowercasing and splitting into words
func CleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}

func main() {
	fmt.Println("Gator Application Start.")
	config, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	s := State{
		config: config,
	}
	fmt.Printf("s: %v\n", s)

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	words := CleanInput(input)
	command := words[0]
	fmt.Printf("command: %s\n", command)
	args := words[1:]
	fmt.Printf("args: %s\n", args)
}
