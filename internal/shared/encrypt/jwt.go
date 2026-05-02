package encrypt

import (
	"os"
	"strconv"
	"time"

	"github.com/Zyprush18/badmintonzz/internal/shared/errs"
	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWToken(user_id int, email, role string) (string, error) {
	claims := customClaims{
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: strconv.Itoa(user_id),
			Subject: email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}


func ParseJWToken(token string) (*customClaims, error) {
	t, err := jwt.ParseWithClaims(token, &customClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	
	if err != nil {
		return nil, err
	}

	claims, ok:= t.Claims.(*customClaims)
	if !ok {
		return nil, errs.InvalidClaims
	}

	return claims, nil
}