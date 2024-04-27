package main

import (
	"Go-Project-Portfolio-RestfulAPI/api"
	"Go-Project-Portfolio-RestfulAPI/postgres"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	var db = postgres.NewGormPostgres()
	defer func() {
		dbPointer, err := db.DB()
		if err != nil {
			fmt.Println(err)
		}
		err2 := dbPointer.Close()
		_ = err2
	}()

	postgres.Migrate(db)

	server := echo.New()
	api.SetRouting(server)
	server.Start("localhost:8080")
}
