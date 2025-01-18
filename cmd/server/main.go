package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/luckygoyal039/youtube-fetcher/internal/config"
	"github.com/luckygoyal039/youtube-fetcher/internal/handler"
	"github.com/luckygoyal039/youtube-fetcher/internal/models"
	"github.com/luckygoyal039/youtube-fetcher/internal/repository"
	"github.com/luckygoyal039/youtube-fetcher/internal/service"
	"github.com/luckygoyal039/youtube-fetcher/internal/worker"
	"github.com/luckygoyal039/youtube-fetcher/pkg/youtube"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()

	// Initialize database
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Video{})

	// Initialize components
	youtubeClient := youtube.NewClient(cfg.YouTubeAPIKeys)
	videoRepo := repository.NewVideoRepository(db)
	youtubeService := service.NewYouTubeService(youtubeClient, videoRepo)
	videoHandler := handler.NewVideoHandler(videoRepo)

	// Start background worker
	fetcher := worker.NewVideoFetcher(youtubeService, cfg.SearchQuery, cfg.FetchInterval)
	fetcher.Start()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(logger.New())  // Request logging
	app.Use(recover.New()) // Panic recovery

	// Routes
	api := app.Group("/api")
	api.Get("/videos", videoHandler.ListVideos)

	log.Fatal(app.Listen(":8080"))
}
