package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", NewError("failed to generate password", ErrInternal)
	}
	return string(pass), nil
}
func ComparePass(hashed, current string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(current)); err != nil {
		return NewError("wrong password provided", ErrBadRequest)
	}
	return nil
}
