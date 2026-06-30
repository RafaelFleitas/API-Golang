package model

import (
	"os"
	"strings"
	"time"

	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/golang-jwt/jwt"
)

var JWT_SECRET_KEY = "JWT_SECRET_KEY"

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {

	secret := os.Getenv(JWT_SECRET_KEY)

	//Informações que serão exibidas
	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), //tempo de expiração do token
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //Cria o token JWT com as informações de login

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError("Error generating token")
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")

	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	return &userDomain{
		id:    int64(claims["id"].(float64)),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil

}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer: ") {
		token = strings.TrimPrefix(token, "Bearer: ")

		return token
	}

	return token

}
