package response

import (
	vo "training-plan/internal/api/domain/plan/entities"
	"training-plan/internal/api/domain/user_activity"
)

type GetUserProfileResponse struct {
	Username               string                     `json:"username"`
	JoinedDate             string                     `json:"joined_date"`
	LatestRunningProfile   FindRunningProfileResponse `json:"latest_running_profile"`
	MostCommonActivityType vo.ActivityType            `json:"most_common_activity_type"`
	LongestRun             useractivity.ActivityStats `json:"longest_run"`
}
