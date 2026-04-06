package encrypt

import (
	"os"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWToken(user_id int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}


func ParseJWToken(token string) (int, string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	
	if err != nil {
		return 0, "", err
	}
	// if !t.Valid {
	// 	return 0, "", fmt.Errorf("invalid token")
	// }

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errs.InvalidClaims
	}
	
	user_id := int(claims["user_id"].(float64))
	email := claims["email"].(string)
	return user_id, email, nil
}