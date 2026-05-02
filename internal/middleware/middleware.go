package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Zyprush18/badmintonzz/internal/shared/encrypt"
	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/gin-gonic/gin"
)


func CheckAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": errs.NoAuthorizationHeader})
			c.Abort()
			return
		}

		token := strings.Split(authHeader, " ")
		if len(token) < 2 || token[0] != "Bearer" || token[1] == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization header"})
			c.Abort()
			return
		}


		claims, err := encrypt.ParseJWToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		id_user, err := strconv.Atoi(claims.ID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid user ID"})
			c.Abort()
			return
		}

		c.Set("user_id", id_user)
		c.Set("email", claims.Subject)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func CheckRole(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_role:= c.GetString("role")

		for _, r := range role {
			if user_role == r {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": errs.ForbiddenAccess})
		c.Abort()
	}
}