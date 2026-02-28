package model

import "time"

type Staff struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Username     string `gorm:"uniqueIndex:ux_staff_login;not null"`
	PasswordHash string `gorm:"not null"`

	HospitalID uint `gorm:"uniqueIndex:ux_staff_login;not null;index"`
	// (optional) Hospital Hospital `gorm:"foreignKey:HospitalID"`
}
