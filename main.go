package main

import (
	"github.com/MaxSkold/projectMessenger/internal/config"
	"log"
)

func main() {
	dbconfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("Error loading DB config: %v", err)
	}

	db, err := dbconfig.Connect()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	log.Println("App is running")
}
