package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/goIdioms/store/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configPath := "config.yaml"
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	db, err := initDatabase()
	// test
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupRoutes(app, db)

	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	log.Printf("Сервер запущен на порту %s", port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown сервера...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Ошибка shutdown сервера: %v", err)
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
