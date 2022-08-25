package db

import (
	// h "github.com/ijasmoopan/usermanagement-api/common/helpers"
	u "github.com/ijasmoopan/usermanagement-api/api/users"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func ConnectDB() *gorm.DB {

	dbURL := "postgres://postgres:ijasmoopan@localhost:5432/myusers"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&u.Users)

	return db
}

