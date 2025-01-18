package config

import (
	"os"
	"strings"
	"time"
)

type Config struct {
	YouTubeAPIKeys []string
	SearchQuery    string
	FetchInterval  time.Duration
	DatabaseURL    string
}

func Load() *Config {
	return &Config{
		YouTubeAPIKeys: strings.Split(os.Getenv("YOUTUBE_API_KEYS"), ","),
		SearchQuery:    os.Getenv("SEARCH_QUERY"),
		FetchInterval:  time.Duration(10) * time.Second,
		DatabaseURL:    os.Getenv("DATABASE_URL"),
	}
}
