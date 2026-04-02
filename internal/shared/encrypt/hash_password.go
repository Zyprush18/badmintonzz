package encrypt

import "golang.org/x/crypto/bcrypt"

func HashingPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}