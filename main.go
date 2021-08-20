package main

// By: DARTHxIKE

import (
	"fmt"

	"github.com/IgorLomba/API-REST-GO/db"
	"github.com/IgorLomba/API-REST-GO/models"
	"github.com/IgorLomba/API-REST-GO/server"
)

func main() {
	db := db.ConnectDb()
	// log.Println(db.RowsAffected)
	fmt.Println("RUNNING..")
	// create tables and migrations
	db.AutoMigrate(models.Person{}, models.Address{})
	server := server.NewServer()
	server.Run()
}
