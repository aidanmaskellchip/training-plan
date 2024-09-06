package valueobjects

type ActivityType struct {
	Type string `json:"type"`
}

var (
	EasyRun      = ActivityType{Type: "easy_run"}
	LongRun      = ActivityType{Type: "long_run"}
	IntervalsRun = ActivityType{Type: "intervals_run"}
	RacePaceRun  = ActivityType{Type: "race_pace_run"}
	GoalRun      = ActivityType{Type: "goal_run"}
)

func GetActivityTypeStrings() []string {
	return []string{
		EasyRun.Type,
		LongRun.Type,
		IntervalsRun.Type,
		RacePaceRun.Type,
		GoalRun.Type,
	}
}

func (at ActivityType) String() string {
	return at.Type
}
