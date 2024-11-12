package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	return string(bytes), err
}

func ValidatePassword(pwd string, hashedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	return err == nil
}
