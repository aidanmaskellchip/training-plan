package model

type ActivityStats struct {
	UserID   string  `json:"user_id"`
	Pace     float32 `json:"pace"`
	Distance float32 `json:"distance"`
}
