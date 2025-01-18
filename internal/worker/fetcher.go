package worker

import (
	"fmt"
	"time"

	"github.com/luckygoyal039/youtube-fetcher/internal/service"
)

type VideoFetcher struct {
	service     *service.YouTubeService
	searchQuery string
	interval    time.Duration
}

func NewVideoFetcher(service *service.YouTubeService, searchQuery string, interval time.Duration) *VideoFetcher {
	return &VideoFetcher{
		service:     service,
		searchQuery: searchQuery,
		interval:    interval,
	}
}

func (f *VideoFetcher) Start() {
	ticker := time.NewTicker(f.interval)
	go func() {
		for range ticker.C {
			err := f.service.FetchAndStoreVideos(f.searchQuery)
			if err != nil {
				fmt.Println("err:", err)
			}
		}
	}()
}
