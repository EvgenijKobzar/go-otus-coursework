package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func MakeByLogin(login string) (string, error) {
	// Секрет из .env
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Время истечения
	expireHours, _ := time.ParseDuration(fmt.Sprintf("%sh", os.Getenv("JWT_EXPIRE_HOURS")))

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(expireHours).Unix(),
	})

	// Подписываем токен
	return token.SignedString(jwtSecret)
}
