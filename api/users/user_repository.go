package users

import (
	// "database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	// Id        int `json:"id"`
	Id int `json:"id" gorm:"type:bigserial;primaryKey;autoIncrement"`
	// Username  string `json:"username" gorm:"uniqueIndex:idx_username;not null"`
	Username  string     `json:"username" gorm:"type:varchar(255);not null"`
	Password  string     `json:"password" gorm:"type:varchar(255);not null"`
	Status    bool       `json:"status"  gorm:"type:bool;default:true"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:now()"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"default:NULL"`
	// DeletedAt *time.Time `json:"deleted_at" gorm:"default:NULL;DeletedAt"`
	DeletedAt gorm.DeletedAt `json: "deleted_at"`
}

var Users = []User{}

type FormData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
