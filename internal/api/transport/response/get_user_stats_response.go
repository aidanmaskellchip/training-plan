package response

import (
	"training-plan/internal/api/domain/user_activity"
)

type GetUserStatsResponse struct {
	UserFastestRun      useractivity.ActivityStats `json:"user_fastest_run"`
	CommunityFastestRun useractivity.ActivityStats `json:"community_fastest_run"`
	UserLongestRun      useractivity.ActivityStats `json:"user_longest_run"`
	CommunityLongestRun useractivity.ActivityStats `json:"community_longest_run"`
}
