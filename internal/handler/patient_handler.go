package handler

import (
	"net/http"

	"github.com/OrioXZ/agnos-backend-assignment/internal/dto"
	"github.com/OrioXZ/agnos-backend-assignment/internal/service"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	svc *service.PatientService
}

func NewPatientHandler(svc *service.PatientService) *PatientHandler {
	return &PatientHandler{svc: svc}
}

func (h *PatientHandler) SearchByID(c *gin.Context) {
	id := c.Param("id")
	hid := c.GetUint("hospitalId")

	out, err := h.svc.SearchByID(hid, id) // service method ใหม่
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "patient not found"})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *PatientHandler) Search(c *gin.Context) {
	var in dto.PatientSearchInput
	if err := c.ShouldBindQuery(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid query"})
		return
	}

	hid := c.GetUint("hospitalId") // มาจาก AuthJWT middleware
	out, err := h.svc.Search(hid, in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "system error"})
		return
	}
	c.JSON(http.StatusOK, out)
}
