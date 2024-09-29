package main

import (
	"fmt"
	"github.com/kidmuon/gator_bootdev/internal/config"
	"log"
	"os"
)

func main() {
	programConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	programState := &State{&programConfig}

	programCommands := Commands{validCommands: make(map[string]func(*State, Command) error)}
	programCommands.register("login", handlerLogin)

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
