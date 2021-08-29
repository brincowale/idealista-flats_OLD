package storage

import (
	"fmt"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"idealista-flats/idealista"
)

var DB *gorm.DB

type Deal struct {
	Id string
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("flats.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&idealista.Property{})
}

func InsertDeal(property idealista.Property) bool {
	property.Id = fmt.Sprintf("%s_%s_%.0f", property.Address, property.Floor, property.Price)
	err := DB.Create(property)
	return err.Error == nil
}