package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	ShortCode string `gorm:"uniqueIndex;not null"`
	LongURL   string `gorm:"not null"`
}
