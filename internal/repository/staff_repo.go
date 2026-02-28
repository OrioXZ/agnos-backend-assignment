package repository

import (
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"gorm.io/gorm"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{db: db}
}

func (r *StaffRepository) Create(staff *model.Staff) error {
	return r.db.Create(staff).Error
}

func (r *StaffRepository) FindByUsernameAndHospitalID(username string, hospitalID uint) (*model.Staff, error) {
	var s model.Staff
	if err := r.db.Where("username = ? AND hospital_id = ?", username, hospitalID).First(&s).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
