package handler

import (
	"net/http"

	"github.com/OrioXZ/agnos-backend-assignment/internal/service/auth"
	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	staffService *auth.StaffService
}

func NewStaffHandler(staffService *auth.StaffService) *StaffHandler {
	return &StaffHandler{staffService: staffService}
}

type createStaffRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	HospitalCode string `json:"hospitalCode"`
}

func (h *StaffHandler) Create(c *gin.Context) {
	var req createStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	staff, err := h.staffService.CreateStaff(auth.CreateStaffInput{
		Username:     req.Username,
		Password:     req.Password,
		HospitalCode: req.HospitalCode,
	})
	if err != nil {
		// แยก status แบบง่ายก่อน
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         staff.ID,
		"username":   staff.Username,
		"hospitalId": staff.HospitalID,
	})
}
