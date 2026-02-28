package repository

import (
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"gorm.io/gorm"
)

type HospitalRepository struct {
	db *gorm.DB
}

func NewHospitalRepository(db *gorm.DB) *HospitalRepository {
	return &HospitalRepository{db: db}
}

func (r *HospitalRepository) FindByCode(code string) (*model.Hospital, error) {
	var h model.Hospital
	if err := r.db.Where("code = ?", code).First(&h).Error; err != nil {
		return nil, err
	}
	return &h, nil
}
