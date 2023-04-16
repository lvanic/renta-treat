package utils

import (
	"go/pkg/model"
	"regexp"
)

func ValidateUser(user *model.User) (bool, string) {
	if user.Age < 18 {
		return false, "Age must be over 18"
	} else if match, _ := regexp.MatchString(`^([a-z0-9_-]+.)*[a-z0-9_-]+@[a-z0-9_-]+(.[a-z0-9_-]+)*.[a-z]{2,6}$`, user.Email); match != true {
		return false, "Email not valid"
	} else {
		return true, "User is valid"
	}
}
