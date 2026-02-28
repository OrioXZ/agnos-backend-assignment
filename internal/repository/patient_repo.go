package repository

import (
	"github.com/OrioXZ/agnos-backend-assignment/internal/dto"
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

func (r *PatientRepository) Search(hospitalID uint, input dto.PatientSearchInput) ([]model.Patient, error) {
	query := r.db.Where("hospital_id = ?", hospitalID)

	if input.NationalID != "" {
		query = query.Where("national_id = ?", input.NationalID)
	}
	if input.PassportID != "" {
		query = query.Where("passport_id = ?", input.PassportID)
	}
	if input.FirstName != "" {
		query = query.Where("first_name ILIKE ? OR first_name ILIKE ?", "%"+input.FirstName+"%", "%"+input.FirstName+"%")
	}
	if input.LastName != "" {
		query = query.Where("last_name ILIKE ? OR last_name ILIKE ?", "%"+input.LastName+"%", "%"+input.LastName+"%")
	}
	if input.DateOfBirth != "" {
		query = query.Where("date_of_birth = ?", input.DateOfBirth)
	}
	if input.PhoneNumber != "" {
		query = query.Where("phone_number = ?", input.PhoneNumber)
	}
	if input.Email != "" {
		query = query.Where("email = ?", input.Email)
	}

	var patients []model.Patient
	err := query.Find(&patients).Error
	return patients, err
}
