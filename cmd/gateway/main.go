package main

import (
	"github.com/MaxSkold/projectMessenger/internal/auth"
	"github.com/MaxSkold/projectMessenger/internal/logger"
	"github.com/MaxSkold/projectMessenger/internal/server"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	defer logger.InitLogger()()

	db, err := connectDB()
	if err != nil {
		logger.Log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	r := router.New()

	// Initialize the application
	handlerAuth := server.StartAuthServer(db)
	auth.RegisterAuthRoutes(r, handlerAuth)

	logger.Log.Infow("🌐 Starting server...", "port", 8080)
	if err := fasthttp.ListenAndServe(":8080", corsMiddleware(r.Handler)); err != nil {
		logger.Log.Fatalw("Server failed", "error", err)
	}

	logger.Log.Info("🛑 Server shut down")
}
