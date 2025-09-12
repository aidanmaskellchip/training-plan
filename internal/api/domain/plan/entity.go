package plan

import (
	"github.com/google/uuid"
	"training-plan/internal/api/domain/plan/entities"
)

type Entity struct {
	ID           uuid.UUID             `json:"id"`
	UserID       uuid.UUID             `json:"user_id"`
	Length       int                   `json:"length"`
	GoalDistance string                `json:"goal_distance"`
	Week1        entities.ActivityWeek `json:"week_1"`
	Week2        entities.ActivityWeek `json:"week_2"`
	Week3        entities.ActivityWeek `json:"week_3"`
	Week4        entities.ActivityWeek `json:"week_4"`
	Week5        entities.ActivityWeek `json:"week_5"`
	Week6        entities.ActivityWeek `json:"week_6"`
	Week7        entities.ActivityWeek `json:"week_7"`
	Week8        entities.ActivityWeek `json:"week_8"`
	Week9        entities.ActivityWeek `json:"week_9"`
	Week10       entities.ActivityWeek `json:"week_10"`
	Week11       entities.ActivityWeek `json:"week_11"`
	Week12       entities.ActivityWeek `json:"week_12"`
}
