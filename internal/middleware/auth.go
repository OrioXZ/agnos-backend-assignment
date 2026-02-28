package middleware

import (
	"net/http"
	"strings"

	"github.com/OrioXZ/agnos-backend-assignment/internal/service/auth"
	"github.com/gin-gonic/gin"
)

const CtxHospitalKey = "hospital"
const CtxUsernameKey = "username"

func AuthJWT(authSvc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(h, "Bearer ")

		claims, err := authSvc.Parse(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		hospital, _ := claims["hospital"].(string)
		username, _ := claims["username"].(string)
		if hospital == "" || username == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		c.Set(CtxHospitalKey, hospital)
		c.Set(CtxUsernameKey, username)
		c.Next()
	}
}
