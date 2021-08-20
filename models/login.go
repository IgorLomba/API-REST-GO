package models

import (
	"log"

	"github.com/IgorLomba/API-REST-GO/db"
	"github.com/IgorLomba/bckp/API-REST-GO/models"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToLogin(login Login) (person models.Person, err error) {
	db := db.ConnectDb()
	err = db.Where("email = ?", login.Email).First(&person).Error
	log.Println(person)
	return
}
