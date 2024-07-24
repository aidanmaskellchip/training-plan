package valueobjects

import "encoding/json"

type RunningDays struct {
	Days []int
}

func NewRunningDays(days []int) *RunningDays {
	return &RunningDays{
		Days: days,
	}
}

func (rd *RunningDays) ToJson() ([]byte, error) {
	days, err := json.Marshal(rd.Days)
	if err != nil {
		return nil, err
	}

	return days, nil
}

func FromJson(days []byte) (RunningDays, error) {
	rd := RunningDays{}
	err := json.Unmarshal(days, &rd.Days)
	if err != nil {
		return rd, err
	}

	return rd, nil
}
