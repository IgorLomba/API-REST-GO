package models

import (
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/db"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(login Login) (err error) {
	db := db.ConnectDb()
	err = db.Where("email=?", person.Email).First(&person).Error
}
