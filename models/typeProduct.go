package models

import "gorm.io/gorm"

type TypeProduct struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
