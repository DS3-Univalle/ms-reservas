package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN = "host=dpg-cl0rv20p2gis739juvi0-a.oregon-postgres.render.com user=userreservas password=VCIhdFuMz0Jzku723KeTjq80zgQeVV2E dbname=reservasdb port=5432"

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
