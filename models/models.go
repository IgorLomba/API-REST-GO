package models

import (
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/db"
	"github.com/jinzhu/gorm"
	"gorm.io/plugin/soft_delete"
)

// struct that the table will be created with
type Person struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Birth   string  `json:"birth"`
	Nif     string  `json:"nif"`
	Address Address `gorm:"foreignKey:id" json:"address"`
	gorm.Model
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

// Many Has One
type Address struct {
	City   string `json:"city"`
	Street string `json:"street"`
	gorm.Model
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

func (Person) TableName() string {
	return "person"
}
func (Address) TableName() string {
	return "address"
}

// i will implement
// func ListDeleted() (person []Person, err error) {
// 	db := db.ConnectDb()
// 	var address []Address

// }

func LoadPeople() (person []Person, err error) {
	var address []Address
	db := db.ConnectDb()

	err = db.Find(&person).Error
	if err == nil {
		err = db.Find(&address).Error
	}

	for i := 0; i < len(person); i++ {
		person[i].Address = address[i]
	}
	return
}

// find and show person by id
func LoadPersonByID(search string) (person Person, err error) {
	db := db.ConnectDb()
	var address Address

	err = db.Find(&person, search).Error
	if err == nil {
		err = db.Find(&address, search).Error
	}

	person.Address = address
	return
}

// search a person by name or part of name
func LoadPersonByName(search string) (person []Person, err error) {
	db := db.ConnectDb()
	var address []Address

	search = "%" + search + "%"

	// you have to use the query builder to compare is_del
	err = db.Table(`person`).Joins(`join address on (person.id = address.id)`).Where(`lower(person.name) like lower(?) and (person.is_del = 0)`, search).Scan(&person).Error
	if err == nil {
		err = db.Table(`address`).Joins(`join person on (person.id = address.id)`).Where(`lower(person.name) like lower(?) and (address.is_del = 0)`, search).Scan(&address).Error
	}

	for i := 0; i < len(person); i++ {
		person[i].Address = address[i]
	}
	return
}

// search a person by your address or part of it address
func LoadPersonByAddress(search string) (person []Person, err error) {
	db := db.ConnectDb()
	var address []Address

	search = "%" + search + "%"

	err = db.Table(`person`).Joins(`join address on (person.id = address.id)`).Where(`lower(address.street) like lower(?) or lower(address.city) like lower(?) and (person.is_del = 0)`, search, search).Scan(&person).Error
	if err == nil {
		err = db.Table(`address`).Joins(`join person on (person.id = address.id)`).Where(`lower(address.street) like lower(?) or lower(address.city) like lower(?) and (address.is_del = 0)`, search, search).Scan(&address).Error
	}

	for i := 0; i < len(person); i++ {
		person[i].Address = address[i]
	}
	return
}

// create a person with address
func CreatePerson(person Person) (Person, error) {
	db := db.ConnectDb()
	err := db.Create(&person).Error
	db.AutoMigrate(Person{}, Address{})
	return person, err
}

// update a person
func UpdatePerson(person Person, address Address) (Person, error) {
	db := db.ConnectDb()
	err := db.Save(&person).Error
	address = person.Address
	if err == nil {
		err = db.Save(&address).Error
	}
	return person, err
}

// find and soft delete person by id
func DeletePersonById(id string) error {
	db := db.ConnectDb()
	err := db.Delete(&Person{}, id).Error
	if err == nil {
		err = db.Delete(&Address{}, id).Error
	}
	return err
}
