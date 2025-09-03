package entities

import "training-plan/internal/api/domain/plan/entities"

type Activity struct {
	Type      string
	Distance  float32
	Pace      float32
	Intervals entities.Intervals
}
