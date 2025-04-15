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
	handler := auth.NewAuthHandler(service)

	r := router.New()
	r.POST("/api/signup", handler.SignUpHandler)

	log.Println("Server is listening :8080 port")
	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
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

func corsMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:5173") // или "*" для всех
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")

		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}

		next(ctx)
	}
}
