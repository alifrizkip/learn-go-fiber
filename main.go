// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"fmt"
	"learn-go-fiber/helper"
	authModule "learn-go-fiber/modules/auth"
	todoModule "learn-go-fiber/modules/todo"
	userModule "learn-go-fiber/modules/user"
	"log"

	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	fmt.Println("Starting server...")

	dsn := "host=localhost user=postgres password=password dbname=go-fiber port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&userModule.User{}, &todoModule.Todo{})

	authServer := authModule.NewServer(db)
	todoServer := todoModule.NewServer(db)

	app := fiber.New()
	app.Use(fiberLogger.New())

	authServer.MountRoutes(app)
	todoServer.MountRoutes(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		response := helper.APIResponse("Not found", fiber.StatusNotFound, false, nil)
		return c.Status(response.Meta.Code).JSON(response)
	})

	log.Fatal(app.Listen(":3000"))
}
