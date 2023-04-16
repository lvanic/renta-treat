package jwt

import (
	"fmt"
	"go/pkg/config"
	"go/pkg/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(user *model.User) (string, error) {
	config := config.NewConfig()
	token_lifespan := config.TokenLifespan
	fmt.Print(token_lifespan)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Secret))
}
