package db

type Feedback struct {
	VideoID  string `json:"video_id"`
	Upvote   uint   `json:"upvote"`
	Downvote uint   `json:"downvote"`
}
