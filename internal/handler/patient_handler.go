package handler

import (
	"net/http"
	"time"

	"github.com/OrioXZ/agnos-backend-assignment/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PatientHandler struct {
	svc *service.PatientService
}

func NewPatientHandler(svc *service.PatientService) *PatientHandler {
	return &PatientHandler{svc: svc}
}

type patientResp struct {
	FirstNameTH  string `json:"first_name_th"`
	MiddleNameTH string `json:"middle_name_th"`
	LastNameTH   string `json:"last_name_th"`

	FirstNameEN  string `json:"first_name_en"`
	MiddleNameEN string `json:"middle_name_en"`
	LastNameEN   string `json:"last_name_en"`

	DateOfBirth *string `json:"date_of_birth"` // yyyy-mm-dd
	PatientHN   string  `json:"patient_hn"`
	NationalID  string  `json:"national_id"`
	PassportID  string  `json:"passport_id"`
	PhoneNumber string  `json:"phone_number"`
	Email       string  `json:"email"`
	Gender      string  `json:"gender"` // M/F
}

func (h *PatientHandler) Search(c *gin.Context) {
	id := c.Param("id")

	// สมมติ middleware ใส่ hospitalId มาแล้ว
	hospitalIDAny, ok := c.Get("hospitalId")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing hospital context"})
		return
	}
	hospitalID := hospitalIDAny.(uint)

	p, err := h.svc.Search(hospitalID, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "patient not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "system error"})
		return
	}

	var dob *string
	if p.DateOfBirth != nil {
		s := p.DateOfBirth.Format("2006-01-02")
		dob = &s
	}

	c.JSON(http.StatusOK, patientResp{
		FirstNameTH:  p.FirstNameTH,
		MiddleNameTH: p.MiddleNameTH,
		LastNameTH:   p.LastNameTH,
		FirstNameEN:  p.FirstNameEN,
		MiddleNameEN: p.MiddleNameEN,
		LastNameEN:   p.LastNameEN,
		DateOfBirth:  dob,
		PatientHN:    p.PatientHN,
		NationalID:   p.NationalID,
		PassportID:   p.PassportID,
		PhoneNumber:  p.PhoneNumber,
		Email:        p.Email,
		Gender:       p.Gender,
	})
}

// กัน unused import time ถ้า file อื่นยังไม่ใช้ (ลบได้ถ้าไม่ต้องใช้)
var _ = time.Now
