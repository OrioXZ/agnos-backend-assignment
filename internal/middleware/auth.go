package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/OrioXZ/agnos-backend-assignment/internal/service/auth"
)

func AuthJWT(jwtSvc *auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(h, "Bearer ")
		claims, err := jwtSvc.Parse(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		username, _ := claims["username"].(string)

		// jwt.MapClaims number → float64
		hidFloat, ok := claims["hospitalId"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token hospitalId"})
			return
		}
		hospitalID := uint(hidFloat)

		c.Set("username", username)
		c.Set("hospitalId", hospitalID)
		c.Next()
	}
}
