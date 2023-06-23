package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error
	dsn := "host=lucky.db.elephantsql.com user=zcvdyobu password=4EwrudELryNU4QVo2KUzdAA_BMNw5iK_ dbname=zcvdyobu port=5432 sslmode=disable "
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("FAILED TO CONNECT TO DATABASE")
	}
}