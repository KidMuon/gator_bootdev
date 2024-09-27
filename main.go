package main

import (
	"fmt"
	"github.com/kidmuon/gator_bootdev/internal/config"
	"log"
)

func main() {
	programConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(programConfig)

	err = programConfig.SetUser("KidMuon")
	if err != nil {
		log.Fatal(err)
	}

	programConfig, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	return
}
