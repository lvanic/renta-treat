package utils

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashPassword(password string, salt []byte) string {
	saltedPassword := []byte(password + string(salt))
	hash := sha512.Sum512(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
