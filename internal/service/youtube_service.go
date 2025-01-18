package service

import (
	"time"

	"github.com/luckygoyal039/youtube-fetcher/internal/models"
	"github.com/luckygoyal039/youtube-fetcher/internal/repository"
	"github.com/luckygoyal039/youtube-fetcher/pkg/youtube"
)

type YouTubeService struct {
	client     *youtube.Client
	repository *repository.VideoRepository
}

func NewYouTubeService(client *youtube.Client, repo *repository.VideoRepository) *YouTubeService {
	return &YouTubeService{
		client:     client,
		repository: repo,
	}
}

func (s *YouTubeService) FetchAndStoreVideos(query string) error {
	latestVideo, err := s.repository.GetLatestVideo()
	if err != nil {
		return err
	}

	publishedAfter := time.Now().Add(-time.Hour)
	if latestVideo != nil {
		publishedAfter = latestVideo.PublishedAt
	}

	response, err := s.client.FetchVideos(query, publishedAfter)
	if err != nil {
		return err
	}

	for _, item := range response.Items {
		video := &models.Video{
			VideoID:          item.ID.VideoID,
			Title:            item.Snippet.Title,
			Description:      item.Snippet.Description,
			PublishedAt:      item.Snippet.PublishedAt,
			ThumbnailDefault: item.Snippet.Thumbnails.Default.URL,
			ThumbnailMedium:  item.Snippet.Thumbnails.Medium.URL,
			ThumbnailHigh:    item.Snippet.Thumbnails.High.URL,
		}

		s.repository.CreateVideo(video)
	}

	return nil
}
