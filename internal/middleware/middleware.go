package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

const KeyResponse = "response"
const KeyError = "error"

func Process(c *gin.Context) {

	err, ok := Internalize(c)
	if ok {
		start := time.Now()
		c.Next()
		responseError, _ := c.Get(KeyError)

		if responseError == nil {
			end := time.Now()
			duration := time.Since(start)

			response, _ := c.Get(KeyResponse)

			if !c.Writer.Written() {
				c.JSON(c.Writer.Status(), gin.H{
					"result": response,
					"time": gin.H{
						"start":             start.UnixNano(),
						"end":               end.UnixNano(),
						"duration":          duration,
						"duration_datetime": duration.String(),
						"start_datetime":    start.String(),
						"end_datetime":      end.String(),
					},
				})
			}
		} else {
			err = responseError.(error)
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
}

func Auth(c *gin.Context) {
	// Получаем токен из заголовка
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Проверяем формат "Bearer <token>"
	if len(authHeader) < 8 || authHeader[:7] != "Bearer " {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		return
	}

	tokenString := authHeader[7:]

	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Проверяем claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Извлекаем login
		if _, ok := claims["login"].(string); ok {
			c.Set("isAuthorize", ok)
			c.Next()
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
}
