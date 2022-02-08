package model

import "gorm.io/gorm"

// People define user info
type People struct {
	gorm.Model

	UID      uint64 `gorm:"column:uid;uniqueIndex;not null;"`
	Nickname string `gorm:"index;not null;size:64"`
	Username string `gorm:"uniqueIndex;not null;size:32"` // only english
}
