package main

import (
	"database/sql"
	"fmt"
	"github.com/MaxSkold/projectMessenger/internal/config"
	"github.com/valyala/fasthttp"
	"log"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	_, err := ctx.WriteString("Hello World")
	if err != nil {
		log.Printf("Error in requestHandler: %v", err)
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		} else {
			log.Printf("Database connection closed successfully")
		}
	}()

	server := fasthttp.Server{
		Handler: requestHandler,
	}
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatalf("Error in server: %v", err)
	}

	log.Println("App is running")
}

func connectDB() (*sql.DB, error) {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading DB config: %v", err)
	}

	db, err := dbConfig.Connect()
	if err != nil {
		return nil, fmt.Errorf("critical database connection error: %v", err)
	}

	return db, nil
}
