package config

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadEnv() error {
	err := godotenv.Load("./internal/config/.env")
	if err != nil {
		return errors.New("error loading .env file")
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func NewDBConfig() (*DBConfig, error) {
	err := LoadEnv()
	if err != nil {
		return nil, err
	}

	return &DBConfig{
		Host:     GetEnv("DB_HOST"),
		Port:     GetEnv("DB_PORT"),
		User:     GetEnv("DB_USER"),
		Password: GetEnv("DB_PASSWORD"),
		DBName:   GetEnv("DB_NAME"),
	}, nil
}

func (dbConfig *DBConfig) Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to the database!")
	return db, nil
}
