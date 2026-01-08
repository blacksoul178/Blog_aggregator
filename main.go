package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/blacksoul178/Blog_aggregator/internal/config"
	"github.com/blacksoul178/Blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg_read, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	programState := new(state) // create state
	programState.cfg = cfg_read
	db, err := sql.Open("postgres", programState.cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	dbQueries := database.New(db)
	programState.db = dbQueries

	cmds := new(commands) //initialize and populate commands
	cmds.Commands = make(map[string]func(*state, command) error)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	user_input := os.Args //read input, make sure something was typed
	if len(user_input) < 2 {
		log.Fatal("User input error: Not enough arguments.\n")
	}
	cmd_to_run := command{ // create an object to pass to run
		Name: user_input[1],
		Args: user_input[2:],
	}

	err = cmds.run(programState, cmd_to_run)
	if err != nil {
		log.Fatal(err)
	}
}
