package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/OrioXZ/agnos-backend-assignment/internal/config"
	"github.com/OrioXZ/agnos-backend-assignment/internal/handler"
	"github.com/OrioXZ/agnos-backend-assignment/internal/middleware"
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"github.com/OrioXZ/agnos-backend-assignment/internal/repository"
	"github.com/OrioXZ/agnos-backend-assignment/internal/service"
	"github.com/OrioXZ/agnos-backend-assignment/internal/service/auth"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Quick sanity check
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	// migrate
	if err := db.AutoMigrate(
		&model.Hospital{},
		&model.Staff{},
		&model.Patient{},
	); err != nil {
		log.Fatal(err)
	}

	// seed hospitals (idempotent)
	seedHospitals(db)
	seedPatients(db)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	jwtSvc := auth.New(cfg.JWTSecret)

	hospitalRepo := repository.NewHospitalRepository(db)
	staffRepo := repository.NewStaffRepository(db)
	staffService := auth.NewStaffService(hospitalRepo, staffRepo, jwtSvc)
	staffHandler := handler.NewStaffHandler(staffService)

	r.POST("/staff/create", staffHandler.Create)
	r.POST("/staff/login", staffHandler.Login)

	patientRepo := repository.NewPatientRepository(db)
	patientSvc := service.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientSvc)

	authMW := middleware.AuthJWT(jwtSvc)
	r.GET("/patient/search/:id", authMW, patientHandler.SearchByID)
	r.GET("/patient/search", authMW, patientHandler.Search)
	// ตัวอย่าง route ที่ต้อง auth
	// authMW := middleware.AuthJWT(jwtSvc)
	// r.GET("/patient/search", authMW, patientHandler.Search)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}

}

func seedHospitals(db *gorm.DB) {
	seeds := []model.Hospital{
		{Code: "hospital_a", Name: "Hospital A"},
		{Code: "hospital_b", Name: "Hospital B"},
	}

	for _, h := range seeds {
		var existing model.Hospital
		err := db.Where("code = ?", h.Code).First(&existing).Error
		if err == nil {
			continue
		}
		_ = db.Create(&h).Error
	}

}

func seedPatients(db *gorm.DB) {
	var count int64
	db.Model(&model.Patient{}).Count(&count)
	if count > 0 {
		return
	}

	seeds := []model.Patient{
		{
			HospitalID:  1,
			PatientHN:   "HN001",
			NationalID:  "1234567890123",
			PassportID:  "P1234567",
			FirstNameTH: "สมชาย",
			LastNameTH:  "ใจดี",
			FirstNameEN: "Somchai",
			LastNameEN:  "Jaidee",
			Gender:      "M",
		},
		{
			HospitalID:  1,
			PatientHN:   "HN002",
			NationalID:  "9876543210987",
			PassportID:  "X7654321",
			FirstNameTH: "สมหญิง",
			LastNameTH:  "สุขใจ",
			FirstNameEN: "Somying",
			LastNameEN:  "Sukjai",
			Gender:      "F",
		},
		{
			HospitalID:  2,
			PatientHN:   "HN003",
			NationalID:  "1111111111111",
			PassportID:  "B1111111",
			FirstNameTH: "ทดสอบ",
			LastNameTH:  "โรงพยาบาลB",
			FirstNameEN: "Test",
			LastNameEN:  "HospitalB",
			Gender:      "M",
		},
	}

	db.Create(&seeds)
}
