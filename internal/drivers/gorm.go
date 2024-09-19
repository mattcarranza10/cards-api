package drivers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	gormDB *gorm.DB
)

func connect() {
	var err error
	gormDB, err = gorm.Open(postgres.Open(getDSN()))
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

func getDSN() string {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")

	return fmt.Sprintf("host=%s user=%s  password=%s dbname=%s port=%s", host, user, password, name, port)
}

func GetGormDB() *gorm.DB {
	if gormDB == nil {
		connect()
	}
	return gormDB
}
