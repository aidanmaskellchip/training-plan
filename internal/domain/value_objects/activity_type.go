package valueobjects

type ActivityType struct {
	Type string `json:"type"`
}

var (
	EASY_RUN      = ActivityType{Type: "easy_run"}
	LONG_RUN      = ActivityType{Type: "long_run"}
	INTERVALS_RUN = ActivityType{Type: "intervals_run"}
	RACE_PACE_RUN = ActivityType{Type: "race_pace_run"}
	GOAL_RUN      = ActivityType{Type: "goal_run"}
)

func GetTypeStrings() []string {
	return []string{
		EASY_RUN.Type,
		LONG_RUN.Type,
		INTERVALS_RUN.Type,
		RACE_PACE_RUN.Type,
		GOAL_RUN.Type,
	}
}
