package main

import (
	"fmt"
	"log"

	"github.com/IgorLomba/API-REST-GO/API-REST-GO/db"
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/models"
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/server"
)

func main() {
	db := db.ConnectDb()

	log.Println(db.RowsAffected)

	fmt.Println("RUNNING..")
	db.AutoMigrate(models.Person{}, models.Address{})
	server := server.NewServer()
	server.Run()
}
