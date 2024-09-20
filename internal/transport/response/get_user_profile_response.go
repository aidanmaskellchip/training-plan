package response

import (
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
)

type GetUserProfileResponse struct {
	Username               string                     `json:"username"`
	JoinedDate             string                     `json:"joined_date"`
	LatestRunningProfile   FindRunningProfileResponse `json:"latest_running_profile"`
	MostCommonActivityType vo.ActivityType            `json:"most_common_activity_type"`
	LongestRun             model.ActivityStats        `json:"longest_run"`
}
