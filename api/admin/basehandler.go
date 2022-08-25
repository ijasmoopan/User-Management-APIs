package admin

import "gorm.io/gorm"

type BaseHandler struct {
	db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler {
		db: db,
	}
}