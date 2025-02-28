package databases

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {
	var db *gorm.DB
	var err error

	dbDriver := os.Getenv("DB_DRIVER")
	var dsn string

	if os.Getenv("MYSQL_HOST") == "localhost" {
		os.Setenv("MYSQL_HOST", "127.0.0.1")
	}

	switch dbDriver {
	case "postgres", "psql":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_NAME"), os.Getenv("POSTGRES_PORT"), url.QueryEscape(os.Getenv("TIMEZONE")),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	case "mysql", "mariadb":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
			os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"), url.QueryEscape(os.Getenv("TIMEZONE")),
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	default:
		dsn = fmt.Sprintf("./%s.db", os.Getenv("SQLITE_PATH"))
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Panicln("Failed to connect to database:", err)
		return nil
	}

	log.Println("Successfully connected to:", dbDriver)
	return db
}
