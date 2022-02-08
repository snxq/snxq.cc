package model

import "gorm.io/gorm"

// Article .
type Article struct {
	gorm.Model

	Title   string `gorm:"size:256;not null;uniqueIndex"`
	Author  string `gorm:"index"`
	Summary string `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
}
