package migrations

// By: DARTHxIKE

import (
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/models"
	"github.com/jinzhu/gorm"
)

/* useless function*/
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Person{}, models.Address{})
}
