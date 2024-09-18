package main

import (
	"log"
	"zxsttm/database"
	"zxsttm/server"
	"zxsttm/server/config"
)

func main() {

	// LoadConfig()
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// InitializeDatabase()
	db, err := database.MySQLConnect(&config.MySQL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// StartServer()
	server.StartServer(config.ServePort, db)
}
