# YouTube Fetcher

A Go application that fetches YouTube videos based on search queries and stores them in a PostgreSQL database. Built with Fiber framework for the API endpoints.

## Features

- Periodic YouTube video fetching based on configured search queries
- REST API to retrieve stored videos
- Pagination support for video listing
- Background worker for automatic video fetching
- PostgreSQL database for persistent storage

## Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/Luckygoyal039/youtube-fetcher.git
   cd youtube-fetcher
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   # Database
   DATABASE_URL=postgresql://postgres:password@localhost:5432/youtube_fetcher?sslmode=disable

   # YouTube API
   YOUTUBE_API_KEYS=your_api_key1,your_api_key2

   # Application Config
   SEARCH_QUERY=golang programming
   ```

4. **Set up the database**

   connect to an existing PostgreSQL instance by updating the DATABASE_URL in `.env`

## Running the Application

1. **Start the application**
   ```bash
   go run cmd/server/main.go
   ```

   The application will:
   - Connect to the PostgreSQL database
   - Run database migrations
   - Start the background video fetcher
   - Start the HTTP server on port 8080

2. **Access the API**

   List videos (with pagination):
   ```bash
   curl "http://localhost:8080/api/videos?page=1&page_size=10"
   ```

## API Endpoints

### GET /api/videos
Lists stored videos with pagination support.

Query Parameters:
- `page` (default: 1): Page number
- `page_size` (default: 10): Number of items per page
