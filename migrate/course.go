package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Course struct {
	gorm.Model
	Key      int    `gorm:"column:key"`
	Title    string `gorm:"column:name" json:"name"`
	Semester string `gorm:"column:semester" json:"semester"`
	Grade    string `gorm:"column:grade" json:"grade"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Course{})
	isPresent := db.HasTable(&Course{})
	fmt.Println("Table course is present", isPresent)
	if !isPresent {
		db.CreateTable(&Course{})
	}

	return db
}