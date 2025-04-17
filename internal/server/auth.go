package server

import (
	"github.com/MaxSkold/projectMessenger/internal/auth"
	"gorm.io/gorm"
)

func StartAuthServer(db *gorm.DB) *auth.HandlerAuth {
	//repo := auth.NewMapsCredRepo()
	repo := auth.NewPostgresCredRepo(db)
	service := auth.NewServiceAuth(repo)
	handler := auth.NewAuthHandler(service)

	return handler
}
