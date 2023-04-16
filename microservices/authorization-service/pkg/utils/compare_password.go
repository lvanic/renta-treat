package utils

func ComparePassword(password string, salt []byte, hashedPasswordExpected string) bool {
	if hashedPassword := HashPassword(password, salt); hashedPassword == hashedPasswordExpected {
		return true
	} else {
		return false
	}
}
