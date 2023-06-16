package util

import (
	"fmt"
	"personal/TODO-list/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	WatId string `json:"wat_id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

func CreateJWT(watID, email, name string) (response string, err error) {
	claims := &Claims{
		watID,
		email,
		name,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(setting.JwtSetting.TokenDuration).Unix(),
			Issuer:    setting.AppSetting.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(setting.JwtSetting.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValifyToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(setting.JwtSetting.JwtSecret), nil
	})
}

// GetRandomOAuthStateString will return random string
func GetRandomOAuthStateString() string {
	return setting.JwtSetting.RandOAuthStateString
}
