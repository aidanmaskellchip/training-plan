package response

import "training-plan/internal/domain/model"

type GetUserProfileResponse struct {
	Username               string                     `json:"username"`
	JoinedDate             string                     `json:"joined_date"`
	LatestRunningProfile   FindRunningProfileResponse `json:"latest_running_profile"`
	MostCommonActivityType string                     `json:"most_common_activity_type"`
	LongestRun             model.ActivityStats        `json:"longest_run"`
}
