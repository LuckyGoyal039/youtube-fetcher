package models

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	VideoID          string `gorm:"uniqueIndex"`
	Title            string
	Description      string
	PublishedAt      time.Time `gorm:"index"`
	ThumbnailDefault string
	ThumbnailMedium  string
	ThumbnailHigh    string
}
