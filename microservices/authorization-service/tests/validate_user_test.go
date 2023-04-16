package tests

import (
	"go/pkg/model"
	"go/pkg/utils"
	"testing"
)

func TestValidateUser(t *testing.T) {
	testID := 0
	users := map[int]model.User{}
	users[1] = model.User{
		ID:    0,
		Name:  "Yahor",
		Email: "ivanouski@gmail.com",
		Age:   18,
	}

	users[2] = model.User{
		ID:    0,
		Name:  "Yahor",
		Email: "ivanouski@gma.com",
		Age:   18,
	}

	users[3] = model.User{
		ID:    0,
		Name:  "Yahor",
		Email: "ivanouski@gmail.com",
		Age:   19,
	}

	for _, user := range users {
		if isValid, validMessage := utils.ValidateUser(&user); isValid == false {
			t.Error(validMessage, " in test ", testID)
		}
		testID++
	}
}
