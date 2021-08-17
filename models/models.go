package models

import (
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/db"
	"github.com/jinzhu/gorm"
	"gorm.io/plugin/soft_delete"
)

/*
CREATE TABLE person(
   	id SERIAL primary key,
	name varchar(100),
	birth date,
	nif varchar(9) UNIQUE NOT NULL
);

CREATE TABLE address(
    city varchar(225),
	street varchar(225),
	FOREIGN KEY(person) REFERENCES person(id)
);

// novo

CREATE TABLE person(
   	id SERIAL primary key,
	name varchar(100),
	birth date,
	nif varchar(9) UNIQUE NOT NULL
);

CREATE TABLE address(
	id SERIAL,
    city varchar(225),
	street varchar(225),
	CONSTRAINT fk_person
		FOREIGN KEY(id)
			REFERENCES person(id)
);

insert into person(name,birth,nif)
values('test1', '10/10/1992', 111111111),
values('test2', '01/01/1991', 222222222);


insert into address(city,street)
values ('linha1', 'linha1'),
values ('linha2', 'linha2');


*/

// gorm.Model faz um soft delete, e etc
type Person struct {
	ID      uint    `json:"id" gorm:"primaryKey"`
	Name    string  `json:"name"`
	Birth   string  `json:"birth"`
	Nif     string  `json:"nif"`
	Address Address `gorm:"foreignKey:id" json:"address"`
	gorm.Model
	IsDel soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

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

// func ListDeleted() (person []Person, err error) {
// 	db := db.ConnectDb()
// 	var address []Address

// }

func LoadPeople() (person []Person, err error) {
	var address []Address
	db := db.ConnectDb()

	/*
		não precisa de nada disso..
		inner join address on (person.id = address.id)
		inner join person on (person.id = address.id)
	*/

	// db.Raw(`select * from person`).Scan(&person)
	// db.Raw(`select * from address`).Scan(&address)

	err = db.Find(&person).Error
	if err == nil {
		err = db.Find(&address).Error
	}

	for i := 0; i < len(person); i++ {
		person[i].Address = address[i]
	}
	return
}

// Só dar um select pelo id e dar append em person
func LoadPersonByID(search string) (person Person, err error) {
	db := db.ConnectDb()
	var address Address

	err = db.Find(&person, search).Error
	if err == nil {
		err = db.Find(&address, search).Error
	}
	// err = db.Raw(`select * from person where person.id = ?`, search).Scan(&person).Error
	// if err == nil {
	// 	err = db.Raw(`select * from address where address.id = ?`, search).Scan(&address).Error
	// }

	person.Address = address
	return
}

func LoadPersonByString(search string) (person []Person, err error) {
	db := db.ConnectDb()
	var address []Address

	// err = db.Find(&person, search).Error
	// if err == nil {
	// 	err = db.Find(&address, search).Error
	// }

	db.Raw(`select * from person join address on (person.id = address.id) where (address.street like ? or address.city like ?)`, search, search).Scan(&person)

	db.Raw(`select * from address where (address.street like ? or address.city like ?)`, search, search).Scan(&address)

	// err = db.Raw(`select * from person where person.id = ?`, search).Scan(&person).Error
	// if err == nil {
	// 	err = db.Raw(`select * from address where address.id = ?`, search).Scan(&address).Error
	// }

	for i := 0; i < len(person); i++ {
		person[i].Address = address[i]
	}
	return
}

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

func LoadPersonByAddress(search string) (person []Person, err error) {
	db := db.ConnectDb()
	var address []Address

	search = "%" + search + "%"

	err = db.Table(`person`).Joins(`join address on (person.id = address.id)`).Where(`lower(address.street) like lower(?) or lower(address.city) like lower(?) and (person.is_del = 0)`, search, search).Scan(&person).Error
	if err == nil {
		err = db.Table(`address`).Joins(`join person on (person.id = address.id)`).Where(`lower(address.street) like lower(?) or lower(address.city) like lower(?) and (address.is_del = 0)`, search, search).Scan(&address).Error
	}

	/*
		err = db.Raw(`select * from person join address on (person.id = address.id) where (lower(address.street) like lower(?) or lower(address.city) like lower(?))`, search, search).Scan(&person).Error
		if err == nil {
			err = db.Raw(`select * from address where (lower(address.street) like lower(?) or lower(address.city) like lower(?))`, search, search).Scan(&address).Error
		}
	*/
	for i := 0; i < len(person); i++ {
		person[i].Address = address[i]
	}
	return
}

func CreatePerson(person Person) (Person, error) {
	db := db.ConnectDb()
	err := db.Create(&person).Error
	db.AutoMigrate(Person{}, Address{})
	return person, err
}

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
