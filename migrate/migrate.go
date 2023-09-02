package main

import (
	"fmt"
	"shiftlab-go-rest-api/db"
	"shiftlab-go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrate")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Workspace{}, &model.Shift{})
}
