package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/OrioXZ/agnos-backend-assignment/internal/config"
	"github.com/OrioXZ/agnos-backend-assignment/internal/handler"
	"github.com/OrioXZ/agnos-backend-assignment/internal/model"
	"github.com/OrioXZ/agnos-backend-assignment/internal/repository"
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

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	hospitalRepo := repository.NewHospitalRepository(db)
	staffRepo := repository.NewStaffRepository(db)
	staffService := auth.NewStaffService(hospitalRepo, staffRepo)
	staffHandler := handler.NewStaffHandler(staffService)

	r.POST("/staff/create", staffHandler.Create)

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
