package main

import (
	"fmt"
	"github.com/MaxSkold/projectMessenger/internal/auth"
	"github.com/MaxSkold/projectMessenger/internal/config"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Initialize the application
	//repo := auth.NewMapsCredRepo()
	repo := auth.NewPostgresCredRepo(db)
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

func connectDB() (*gorm.DB, error) {
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
