package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/terfo1/news/internal/database"
	"github.com/terfo1/news/internal/models"
	"os"
	"strings"
	"time"
)

func AuthRequire(c *fiber.Ctx) error {
	tokenString := c.Cookies("Authorization")
	if tokenString == "" {
		authHeader := c.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: missing token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		var user models.User
		database.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		c.Locals("user", user)
	} else {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	return c.Next()
}
