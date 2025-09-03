package response

import (
	"training-plan/internal/api/domain/model"
	vo "training-plan/internal/api/domain/plan/entities"
)

type GetUserProfileResponse struct {
	Username               string                     `json:"username"`
	JoinedDate             string                     `json:"joined_date"`
	LatestRunningProfile   FindRunningProfileResponse `json:"latest_running_profile"`
	MostCommonActivityType vo.ActivityType            `json:"most_common_activity_type"`
	LongestRun             model.ActivityStats        `json:"longest_run"`
}
