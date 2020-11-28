package middlewares

import (
	"fmt"
	"learn-go-fiber/config"
	"learn-go-fiber/helper"
	"learn-go-fiber/modules/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"gorm.io/gorm"
)

// JwtAuth ...
func JwtAuth(db *gorm.DB) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:     config.JwtSecretKey,
		ErrorHandler:   jwtErrorHandler,
		SuccessHandler: jwtSuccessHandler(db),
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	return helper.SendAPIResponse(c)("Unauthorized", fiber.StatusUnauthorized, false, nil)
}

func jwtSuccessHandler(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := strings.Split(authHeader, " ")[1]

		token, _ := jwt.Parse(tokenString, nil)
		claim, _ := token.Claims.(jwt.MapClaims)
		userID := int(claim["id"].(float64))
		fmt.Println("userID", userID)

		userRepo := user.NewRepository(db)
		user, err := user.NewService(userRepo).GetUserByID(userID)
		if err != nil {
			fmt.Println("err", err.Error())
			return helper.SendAPIResponse(c)("Unauthorized", fiber.StatusUnauthorized, false, nil)
		}

		c.Locals("current_user", user)

		return c.Next()
	}
}
