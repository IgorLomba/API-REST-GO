package db

// By: DARTHxIKE

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/* you must install pgx and gorm to run*/
/* returns a variable that can be used to call database functions */
func ConnectDb() *gorm.DB {
	/* TimeZone=Asia/Shanghai */
	// you should CHANGE it before run!!!
	dsn := "host=127.0.0.1 user=postgres password=123456 dbname=apiTest port=5432 sslmode=disable"
	// https://github.com/jackc/pgx
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}
