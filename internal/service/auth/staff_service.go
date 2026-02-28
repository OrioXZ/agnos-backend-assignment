package auth

import (
	"errors"
	"strings"

	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"github.com/OrioXZ/agnos-backend-assignment/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StaffService struct {
	hospitalRepo *repository.HospitalRepository
	staffRepo    *repository.StaffRepository
	jwtSvc       *Service
}

func NewStaffService(hospitalRepo *repository.HospitalRepository, staffRepo *repository.StaffRepository, jwtSvc *Service) *StaffService {
	return &StaffService{
		hospitalRepo: hospitalRepo,
		staffRepo:    staffRepo,
		jwtSvc:       jwtSvc,
	}
}

type CreateStaffInput struct {
	Username     string
	Password     string
	HospitalCode string
}

func (s *StaffService) CreateStaff(in CreateStaffInput) (*model.Staff, error) {
	in.Username = strings.TrimSpace(in.Username)
	in.Password = strings.TrimSpace(in.Password)
	in.HospitalCode = strings.TrimSpace(in.HospitalCode)

	if in.Username == "" || in.Password == "" || in.HospitalCode == "" {
		return nil, errors.New("missing required fields")
	}

	h, err := s.hospitalRepo.FindByCode(in.HospitalCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("hospital not found")
		}
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	staff := &model.Staff{
		Username:     in.Username,
		PasswordHash: string(hash),
		HospitalID:   h.ID,
	}

	if err := s.staffRepo.Create(staff); err != nil {
		// unique constraint อาจชนได้
		return nil, err
	}

	// ไม่ส่ง hash กลับ
	staff.PasswordHash = ""
	return staff, nil
}

type LoginInput struct {
	Username     string
	Password     string
	HospitalCode string
}

type LoginResult struct {
	Token      string
	ID         uint
	Username   string
	HospitalID uint
}

func (s *StaffService) Login(in LoginInput) (*LoginResult, error) {
	in.Username = strings.TrimSpace(in.Username)
	in.Password = strings.TrimSpace(in.Password)
	in.HospitalCode = strings.TrimSpace(in.HospitalCode)

	if in.Username == "" || in.Password == "" || in.HospitalCode == "" {
		return nil, errors.New("missing required fields")
	}

	h, err := s.hospitalRepo.FindByCode(in.HospitalCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("hospital not found")
		}
		return nil, err
	}

	// ต้องมี method นี้ใน repo (ดูข้อ 2.1)
	staff, err := s.staffRepo.FindByUsernameAndHospitalID(in.Username, h.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.PasswordHash), []byte(in.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.jwtSvc.GenerateToken(staff.Username, staff.HospitalID)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		Token:      token,
		ID:         staff.ID,
		Username:   staff.Username,
		HospitalID: staff.HospitalID,
	}, nil
}
