package forms

import "time"

// Structure Post for Posting details
type Post struct {
	Id        string
	Caption   string
	ImageURL  string
	TIMEstamp time.Duration
}
