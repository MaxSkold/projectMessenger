package main

import (
	"database/sql"
	"fmt"
	"github.com/MaxSkold/projectMessenger/internal/auth"
	"github.com/MaxSkold/projectMessenger/internal/config"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	//db, err := connectDB()
	//if err != nil {
	//	log.Fatalf("Error: %v", err)
	//}
	//defer func() {
	//	if err := db.Close(); err != nil {
	//		log.Fatalf("Error closing database connection: %v", err)
	//	} else {
	//		log.Printf("Database connection closed successfully")
	//	}
	//}()

	// Initialize the application
	repo := auth.NewMapsCredRepo()
	service := auth.NewServiceAuth(repo)
	handler := auth.NewAuthHeader(service)

	r := router.New()
	r.POST("/signup", handler.SignUpHandler)

	log.Println("Server is listening :8080 port")
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		panic(err)
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
