package databases

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func DatabaseInit() *gorm.DB {
	var db *gorm.DB
	var err error

	dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	timezone := os.Getenv("TIMEZONE")

	switch dbDriver {
	case "postgres", "psql":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbName=%s port=%s sslmode=disable TimeZone=%s",
			host, user, password, dbName, port, timezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "mysql", "mariadb":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", user, password, host, port, dbName, timezone)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open(fmt.Sprintf("./%s.db", dbName)), &gorm.Config{})
	}

	if err != nil {
		log.Panicln(err)
		return nil
	}

	log.Println("Success connect to:", dbDriver)
	return db
}
