package model

import "time"

type Hospital struct {
	ID        uint   `gorm:"primaryKey"`
	Code      string `gorm:"size:50;uniqueIndex;not null"`
	Name      string `gorm:"size:120;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
