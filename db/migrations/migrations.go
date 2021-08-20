package migrations

import (
	"github.com/IgorLomba/API-REST-GO/models"
	"github.com/jinzhu/gorm"
)

// By: DARTHxIKE

/* useless function*/
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Person{}, models.Address{})
}
