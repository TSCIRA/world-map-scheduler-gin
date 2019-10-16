package models

import "github.com/jinzhu/gorm"

type Boss struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}