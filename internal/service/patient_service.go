package service

import (
	"github.com/OrioXZ/agnos-backend-assignment/internal/dto"
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"github.com/OrioXZ/agnos-backend-assignment/internal/repository"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

// B) ตามข้อ 4.3: search by optional fields -> return list
func (s *PatientService) Search(hospitalID uint, input dto.PatientSearchInput) ([]model.Patient, error) {
	return s.repo.Search(hospitalID, input)
}

// A) lookup by id (national_id หรือ passport_id) -> return single
func (s *PatientService) SearchByID(hospitalID uint, id string) (*model.Patient, error) {
	return s.repo.FindByNationalOrPassport(hospitalID, id)
}
