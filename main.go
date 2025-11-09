package main

import (
	"database/sql"
	"fmt"
	"gator/internal/cli"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/middleware"
	"gator/internal/models"
	"gator/internal/utils"
	"log"
	"os"

	// import the driver, but you don't use it directly anywhere in your
	// code. The underscore tells Go that you're importing it for
	// its side effects, not because you need to use it.
	_ "github.com/lib/pq"
)

var debug bool

func debugLog(logger *log.Logger, format string, v ...any) {
	if debug {
		logger.Printf(format, v...)
	}
}

func main() {
	debug = false

	logger := log.New(os.Stderr, "gator: ", log.LstdFlags|log.Lshortfile)

	cfg, err := config.Read()
	if err != nil {
		logger.Fatalf("error occurred reading config: %v", err)
	}
	debugLog(logger, "config read successful: %v\n", utils.ToJSON(cfg))

	// Open connection to the database
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		logger.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(db)
	s := &models.State{
		Config: &cfg,
		Db:     dbQueries,
	}
	debugLog(logger, "create state successful. s = %v\n", utils.ToJSON(s))

	cmds := cli.NewCommands()
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", middleware.LoggedIn(cli.HandlerAddFeed))
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", middleware.LoggedIn(cli.HandlerFollow))
	cmds.Register("following", middleware.LoggedIn(cli.HandlerFollowing))

	if len(os.Args) < 2 {
		logger.Fatalf("Usage: gator command <arguments>")
		return
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	debugLog(logger, "create command successful. cmd = %s\n", utils.ToJSON(cmd))
	fmt.Printf("run command: %s, args: %s\n", cmd.Name, cmd.Args)

	err = cmds.Run(s, cmd)
	if err != nil {
		logger.Fatalf("error running command: %v\n", err)
	}
}
