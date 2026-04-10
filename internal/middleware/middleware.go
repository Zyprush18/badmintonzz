package middleware

import (
	"net/http"
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

		token := strings.Split(authHeader, ",")
		if len(token) < 2 || token[0] != "Bearer" || token[1] == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": errs.InvalidRequest})
			c.Abort()
			return
		}


		user_id, email , err := encrypt.ParseJWToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Set("user_id", user_id)
		c.Set("email", email)


		c.Next()
	}
}

func CheckRole(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// user_role:= c.GetString("role")
		user_role:= "admin"

		for _, r := range role {
			if user_role == r {
				c.Set("role", "admin") // hapus ketika sudah integrasikan login dan set role nya di middleware checkauth
				c.Set("user_id", 1) // hapus ketika sudah integrasikan login dan set role nya di middleware checkauth
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": errs.ForbiddenAccess})
		c.Abort()
	}
}