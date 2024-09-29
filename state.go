package main

import (
	"github.com/kidmuon/gator_bootdev/internal/config"
	"github.com/kidmuon/gator_bootdev/internal/database"
)

type State struct {
	cfg *config.Config
	db  *database.Queries
}
