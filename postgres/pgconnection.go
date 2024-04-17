package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/url"
	"sync"
	"time"
)

var once sync.Once

func NewGormPostgres() *gorm.DB {
	var instance *gorm.DB
	once.Do(func() {
		tehranTimezone, _ := time.LoadLocation("Asia/Tehran")

		// Connection configuration
		dsn := &url.URL{
			Scheme:   "AirBird",
			User:     url.UserPassword("docker_user", "docker_user"),
			Host:     "localhost",
			Path:     "",
			RawQuery: "sslmode=disable&timezone=" + tehranTimezone.String(),
		}

		// Convert URL to connection string
		connStr := dsn.String()
		_ = connStr
		connStr2 := "host=localhost user=docker_user password=docker_user dbname=docker_user port=5433 sslmode=disable TimeZone=Asia/Shanghai"

		db, err := gorm.Open(postgres.Open(connStr2), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		fmt.Println("Successfully connected to the database!")

		instance = db
	})
	return instance
}
