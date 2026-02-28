package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/OrioXZ/agnos-backend-assignment/internal/config"
	"github.com/OrioXZ/agnos-backend-assignment/internal/repository"
)

func main() {
	cfg := config.Load()

	db, err := repository.NewPostgres(cfg.DbDsn)
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

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
