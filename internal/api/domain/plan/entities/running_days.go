package entities

import (
	"encoding/json"
)

type RunningDays struct {
	Days      []int
	Structure map[ActivityType]int
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

func RunningDaysFromJson(days []byte) (RunningDays, error) {
	rd := RunningDays{}
	err := json.Unmarshal(days, &rd.Days)
	if err != nil {
		return rd, err
	}

	return rd, nil
}

// RDStructure how we calculate what activities are scheduled for which days
//
//	var RDStructure = map[ActivityType][]int{
//		EasyRun:      []int{},
//		LongRun:      []int{},
//		RacePaceRun:  []int{},
//		IntervalsRun: []int{},
//	}
func RDStructure(rd RunningDays, longRunDay, rdsPerWeek int) (rds map[ActivityType][]int, err error) {
	rds = make(map[ActivityType][]int)
	rds[LongRun] = append(rds[LongRun], longRunDay)

	filDays := rd.FilteredDays()

	for _, i := range filDays {
		if i != longRunDay {
			rds[EasyRun] = append(rds[EasyRun], i)
		}
	}

	if rdsPerWeek == 2 {
		return rds, nil
	}

	for _, i := range filDays {
		if i != longRunDay || i != rds[EasyRun][0] {
			rds[RacePaceRun] = append(rds[RacePaceRun], i)
		}
	}

	if rdsPerWeek == 3 {
		return rds, nil
	}

	for _, i := range filDays {
		if i != longRunDay || i != rds[EasyRun][0] || i != rds[RacePaceRun][0] {
			rds[EasyRun] = append(rds[EasyRun], i)
		}
	}

	return rds, nil
}

func (rd *RunningDays) FilteredDays() (days []int) {
	for i, day := range rd.Days {
		if day == 1 {
			days = append(days, i)
		}
	}

	return
}
