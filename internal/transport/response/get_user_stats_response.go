package response

import (
	valueobjects "training-plan/internal/domain/model"
)

type GetUserStatsResponse struct {
	UserFastestRun      valueobjects.ActivityStats `json:"user_fastest_run"`
	CommunityFastestRun valueobjects.ActivityStats `json:"community_fastest_run"`
	UserLongestRun      valueobjects.ActivityStats `json:"user_longest_run"`
	CommunityLongestRun valueobjects.ActivityStats `json:"community_longest_run"`
}
