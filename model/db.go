package model

import (
	"gorm.io/gorm"
	"time"
)

type Cache struct {
	gorm.Model
	CacheKey   string `gorm:"unique;not null"`
	CacheValue string
	Expired    *time.Time
}

type Loan struct {
	gorm.Model
}
