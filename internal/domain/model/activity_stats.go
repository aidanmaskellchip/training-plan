package model

type ActivityStats struct {
	Type     string  `json:"type"`
	UserID   string  `json:"user_id"`
	Pace     float32 `json:"pace"`
	Distance float32 `json:"distance"`
}
