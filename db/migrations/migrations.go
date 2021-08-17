package migrations

import (
	"github.com/IgorLomba/API-REST-GO/API-REST-GO/models"
	"github.com/jinzhu/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Person{}, models.Address{})
}
