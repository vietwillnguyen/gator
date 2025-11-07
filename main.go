package main

import (
	"gator/internal/cli"
	"gator/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error occurred reading config: %v", err)		
	}
	fmt.Printf("Config Read Successfully: %v", cfg)		

	// _ := &models.State{Config: &cfg}

	cmds := cli.NewCommands()
	cmds.Register("login", cli.HandlerLogin)

	// ... parse args and run
}
