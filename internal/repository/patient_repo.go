package repository

import (
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) FindByNationalOrPassport(hospitalID uint, id string) (*model.Patient, error) {
	var p model.Patient
	err := r.db.
		Where("hospital_id = ? AND (national_id = ? OR passport_id = ?)", hospitalID, id, id).
		First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
