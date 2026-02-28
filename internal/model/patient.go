package model

import "time"

type Patient struct {
	ID         uint `gorm:"primaryKey"`
	HospitalID uint `gorm:"not null;index"`

	PatientHN  string `gorm:"size:50;index"`
	NationalID string `gorm:"size:50;index"`
	PassportID string `gorm:"size:50;index"`

	FirstNameTH  string `gorm:"size:120"`
	MiddleNameTH string `gorm:"size:120"`
	LastNameTH   string `gorm:"size:120"`

	FirstNameEN  string `gorm:"size:120"`
	MiddleNameEN string `gorm:"size:120"`
	LastNameEN   string `gorm:"size:120"`

	DateOfBirth *time.Time `gorm:"type:date"`
	PhoneNumber string     `gorm:"size:50"`
	Email       string     `gorm:"size:120"`
	Gender      string     `gorm:"size:1"` // M/F

	Source string `gorm:"size:30;not null;default:'MANUAL'"`
	Status string `gorm:"size:1;not null;default:'Y'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
