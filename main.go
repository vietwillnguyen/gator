package main

import (
	"encoding/json"
	"fmt"
	"gator/internal/cli"
	"gator/internal/config"
	"gator/internal/models"
	"log"
	"os"
)

var debug bool

func stringObjectToJSON(v any) string {
	data, _ := json.MarshalIndent(v, "", "  ")
	return string(data)
}

func debugLog(logger *log.Logger, format string, v ...any) {
	if debug {
		logger.Printf(format, v...)
	}
}

func main() {
	// Enable debug logs if env var is set
	debug = os.Getenv("GATOR_DEBUG") == "1"

	logger := log.New(os.Stderr, "gator: ", log.LstdFlags|log.Lshortfile)

	cfg, err := config.Read()
	if err != nil {
		logger.Fatalf("error occurred reading config: %v", err)
	}
	debugLog(logger, "config read successful: %v\n", stringObjectToJSON(cfg))

	s := &models.State{Config: &cfg}
	debugLog(logger, "create state successful. s = %v\n", stringObjectToJSON(s))

	cmds := cli.NewCommands()
	cmds.Register("login", cli.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("Usage: gator command <arguments>")
		return
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	debugLog(logger, "create command successful. cmd = %s\n", stringObjectToJSON(cmd))

	fmt.Printf("Running command: <%s>, args [%s]\n", cmd.Name, cmd.Args)
}
