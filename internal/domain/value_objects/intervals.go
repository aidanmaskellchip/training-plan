package valueobjects

type Intervals struct {
	FastRepPace      float32 `json:"fastRepPace"`
	WarmUpDistance   float32 `json:"warmUpDistance"`
	FastRepDistance  float32 `json:"fastRepDistance"`
	RecoveryDuration int     `json:"recoveryDuration"`
	NumberOfReps     int     `json:"numberOfReps"`
	WarmDownDistance float32 `json:"warmDownDistance"`
}
