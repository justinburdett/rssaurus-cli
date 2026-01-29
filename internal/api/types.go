package api

import "time"

type Me struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type FeedsResponse struct {
	Feeds []Feed `json:"feeds"`
}

type Feed struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	URL         string     `json:"url"`
	FeedURL     string     `json:"feed_url"`
	FeedType    string     `json:"feed_type"`
	UnreadCount int64      `json:"unread_count"`
	CreatedAt   *time.Time `json:"created_at"`
	LastFetched *time.Time `json:"last_fetched_at"`
}

type ItemsResponse struct {
	Items      []Item `json:"items"`
	NextCursor string `json:"next_cursor"`
}

type Item struct {
	ID          int64      `json:"id"`
	FeedID      int64      `json:"feed_id"`
	FeedTitle   string     `json:"feed_title"`
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	URL         string     `json:"url"`
	PublishedAt *time.Time `json:"published_at"`
	ReadAt      *time.Time `json:"read_at"`
	ImageURL    string     `json:"image_url"`
	AudioURL    string     `json:"audio_url"`
}

type ReadStateResponse struct {
	ID     int64      `json:"id"`
	ReadAt *time.Time `json:"read_at"`
}

type MarkReadResponse struct {
	Updated int64 `json:"updated"`
}

type SavedCreateResponse struct {
	ID        int64      `json:"id"`
	URL       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
}
