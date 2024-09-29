package main

import _ "github.com/lib/pq"

import (
	"database/sql"
	"fmt"
	"github.com/kidmuon/gator_bootdev/internal/config"
	"github.com/kidmuon/gator_bootdev/internal/database"
	"os"
)

func main() {
	programConfig, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", programConfig.DBConnectionString)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	programState := &State{
		cfg: &programConfig,
		db:  dbQueries,
	}

	programCommands := Commands{validCommands: make(map[string]func(*State, Command) error)}
	programCommands.register("login", handlerLogin)
	programCommands.register("register", handlerRegister)
	programCommands.register("reset", handlerReset)
	programCommands.register("users", handlerUsers)
	programCommands.register("agg", handlerAgg)
	programCommands.register("addfeed", handlerAddFeed)

	if len(os.Args) < 2 {
		fmt.Println("error: no command passed")
		os.Exit(1)
	}

	var programArgs []string
	if len(os.Args) == 2 {
		programArgs = []string{}
	} else {
		programArgs = os.Args[2:]
	}

	err = programCommands.run(programState, Command{name: os.Args[1], args: programArgs})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
