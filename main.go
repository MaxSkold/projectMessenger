package main

import (
	"database/sql"
	"fmt"
	"github.com/MaxSkold/projectMessenger/internal/config"
	"log"
)

func connectDB() (*sql.DB, error) {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading DB config: %v", err)
	}

	db, err := dbConfig.Connect()
	if err != nil {
		return nil, fmt.Errorf("critical database connection error: %v", err)
	}

	// Возвращаем базу данных для дальнейшего использования
	return db, nil
}

func main() {
	// Подключаемся к базе данных
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Используем defer для безопасного закрытия соединения
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		} else {
			log.Printf("Database connection closed successfully")
		}
	}()

	log.Println("App is running")
}
