package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func DoLogin(login Login) (err error) {
// 	db := db.ConnectDb()
// 	var person models.Person
// 	err = db.Where("email=?", person.Email).First(&person).Error
// }
