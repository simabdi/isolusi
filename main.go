package main

import (
	"isolusi/internal/config"
	"isolusi/internal/database"
	"isolusi/internal/route"
)

func main() {
	config.Initialize()
	database.Migrate()
	route.Apis()
}
