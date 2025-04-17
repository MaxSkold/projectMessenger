package main

import (
	"fmt"
	"github.com/MaxSkold/projectMessenger/internal/config"
	"github.com/MaxSkold/projectMessenger/internal/logger"
	"github.com/valyala/fasthttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading DB config: %v", err)
	}

	db, err := gorm.Open(postgres.Open(dbConfig.DSN()), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	logger.Log.Infow("Connected to database",
		"host", dbConfig.Host,
		"port", dbConfig.Port,
	)
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
