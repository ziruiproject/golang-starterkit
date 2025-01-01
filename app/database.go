package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var (
	dbInstance *sqlx.DB
	once       sync.Once
)

func ConnectDatabase() (*sqlx.DB, error) {
	var err error
	once.Do(func() {
		if err = godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %v", err)
		}

		requiredEnvVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
		for _, envVar := range requiredEnvVars {
			if os.Getenv(envVar) == "" {
				err = fmt.Errorf("environment variable %s is not set", envVar)
				return
			}
		}

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		)

		dbInstance, err = sqlx.Open("postgres", dsn)
		if err != nil {
			err = fmt.Errorf("failed to open database connection: %v", err)
			return
		}

		if err = dbInstance.Ping(); err != nil {
			err = fmt.Errorf("failed to connect to the database: %v", err)
			return
		}

		log.Println("Successfully connected to the database")
	})
	return dbInstance, err
}

func CloseDatabase() {
	if dbInstance != nil {
		if err := dbInstance.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully")
		}
	}
}
