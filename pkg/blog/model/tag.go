package model

import "gorm.io/gorm"

// Tag label
type Tag struct {
	gorm.Model

	Name string `gorm:"index"`
}
