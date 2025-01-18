package repository

import (
	"github.com/luckygoyal039/youtube-fetcher/internal/models"
	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

func (r *VideoRepository) CreateVideo(video *models.Video) error {
	return r.db.Create(video).Error
}

func (r *VideoRepository) GetLatestVideo() (*models.Video, error) {
	var video models.Video
	err := r.db.Order("published_at desc").First(&video).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &video, err
}

func (r *VideoRepository) ListVideos(page, pageSize int) ([]models.Video, int64, error) {
	var videos []models.Video
	var total int64

	r.db.Model(&models.Video{}).Count(&total)

	err := r.db.Order("published_at desc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&videos).Error

	return videos, total, err
}
