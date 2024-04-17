package main

import (
	"Go-Project-Portfolio-RestfulAPI/postgres"
	"fmt"
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
}
