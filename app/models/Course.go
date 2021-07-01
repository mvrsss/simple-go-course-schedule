package models

import (
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

func (c *Course) TableName() string {
	return "course"
}

func AddCourse(db *gorm.DB, c *Course) (err error) {
	if err = db.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func GetAllCourses(db *gorm.DB, c *[]Course) (err error) {
	if err = db.Order("id desc").Find(c).Error; err != nil {
		return err
	}
	return nil
}

func GetACourse(db *gorm.DB, id int, c *Course) (err error) {
	if err = db.Where("key = ?", id).First(&c).Error; err != nil {
		return err
	}
	return nil
}

func RetakeCourse(db *gorm.DB, c *Course) (err error) {
	if err = db.Save(c).Error; err != nil {
		return err
	}
	return nil
}

func DropCourse(db *gorm.DB, c *Course) (err error) {
	if err = db.Delete(c).Error; err != nil {
		return err
	}
	return nil
}
