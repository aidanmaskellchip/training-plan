package response

import (
	"training-plan/internal/api/domain/model"
)

type GetUserStatsResponse struct {
	UserFastestRun      model.ActivityStats `json:"user_fastest_run"`
	CommunityFastestRun model.ActivityStats `json:"community_fastest_run"`
	UserLongestRun      model.ActivityStats `json:"user_longest_run"`
	CommunityLongestRun model.ActivityStats `json:"community_longest_run"`
}
