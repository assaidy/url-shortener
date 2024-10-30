package server

import (
	"github.com/assaidy/url-shortener/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberServer struct {
	*fiber.App
	DB *database.DBService
}

func NewFiberServer() *FiberServer {
	fs := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "url-shortner",
			AppName:      "url-shortner",
		}),
		DB: database.NewDBService(),
	}
	fs.Use(logger.New())
	return fs
}
