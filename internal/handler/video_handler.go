package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/luckygoyal039/youtube-fetcher/internal/repository"
)

type VideoHandler struct {
	repo *repository.VideoRepository
}

func NewVideoHandler(repo *repository.VideoRepository) *VideoHandler {
	return &VideoHandler{repo: repo}
}

func (h *VideoHandler) ListVideos(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page number",
		})
	}

	pageSize, err := strconv.Atoi(c.Query("page_size", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid page size",
		})
	}

	videos, total, err := h.repo.ListVideos(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":      videos,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}
