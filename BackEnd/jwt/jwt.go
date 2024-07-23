package jwt

import (
	"BackEnd/dto"
	"crypto/rand"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

var Secreto string

func GenerateSecretJWT() (string, error) {
	bytes := make([]byte, 64)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GenerateUserToken(userDto dto.UserDto) (string, error) {
	claims := jwt.MapClaims{
		"id":         userDto.Id,
		"name":       userDto.FirstName,
		"last_name":  userDto.LastName,
		"user_email": userDto.UserEmail,
		"tipo":       userDto.Tipo,
		"dni":        userDto.DNI,
		"telefono":   userDto.Telefono,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(Secreto))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secreto), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func UpdateSecretPeriodically() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	var err error

	for {
		select {
		case <-ticker.C:
			Secreto, err = GenerateSecretJWT() // 32 bytes secret
			if err != nil {
				log.Fatalf("Error generating secret: %v", err)
			}
		}
	}
}
