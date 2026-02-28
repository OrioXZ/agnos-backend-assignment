package service

import (
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"github.com/OrioXZ/agnos-backend-assignment/internal/repository"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) Search(hospitalID uint, id string) (*model.Patient, error) {
	return s.repo.FindByNationalOrPassport(hospitalID, id)
}
