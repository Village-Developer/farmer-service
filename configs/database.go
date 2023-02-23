package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configs struct{}

func (h *Configs) Connect() *gorm.DB {
	// connect to database
	dsn := "root:village@tcp(localhost:3307)/api_document?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:village@tcp(farmerproject-mysql-1:3306)/api_document?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
