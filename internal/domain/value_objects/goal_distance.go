package valueobjects

type GoalDistance struct {
	Type string `json:"goal_distance"`
}

var (
	FiveK        = ActivityType{Type: "five_k"}
	TenK         = ActivityType{Type: "ten_k"}
	HalfMarathon = ActivityType{Type: "half_marathon"}
	FullMarathon = ActivityType{Type: "full_marathon"}
)

func GetTypeStrings() []string {
	return []string{
		FiveK.Type,
		TenK.Type,
		HalfMarathon.Type,
		FullMarathon.Type,
	}
}

func (gd GoalDistance) String() string {
	return gd.Type
}
