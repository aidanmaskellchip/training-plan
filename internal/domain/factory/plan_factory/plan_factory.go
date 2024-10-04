package planfactory

import (
	"sync"
	weekfactory "training-plan/internal/domain/factory/week_factory"
	"training-plan/internal/domain/model"
	valueobjects "training-plan/internal/domain/value_objects"
)

func NewPlan(rp model.RunningProfile) (p model.Plan, err error) {
	p.UserID = rp.UserID
	p.Length = rp.PlanLength
	p.GoalDistance = valueobjects.GoalDistance{Type: rp.GoalDistance}

	var weeks []model.ActivityWeek
	for i := 0; i < rp.PlanLength; i++ {
		weeks = append(weeks, weekfactory.NewWeek())
	}

	rd, err := valueobjects.RunningDaysFromJson(rp.RunningDays)
	if err != nil {
		return model.Plan{}, err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	errChn := make(chan error, 2)

	go func() {
		defer wg.Done()

		err = setLongRuns(&weeks, rp.LongRunDay, rp.PlanLength)
		if err != nil {
			errChn <- err
		}
	}()

	go func() {
		defer wg.Done()

		err = setEasyRuns(&weeks, rd, rp.LongRunDay, rp.PlanLength)
		if err != nil {
			errChn <- err
		}
	}()

	wg.Wait()

	// long runs distances
	// easy run distances
	// threshold

	return p, nil
}

func setLongRuns(weeks *[]model.ActivityWeek, longRunDay int, planLength int) error {
	for i, w := range *weeks {
		if i == len(*weeks)-1 {
			return nil
		}

		lrd, err := w.GetDayByIndex(longRunDay)
		if err != nil {
			return err
		}

		lrd.Distance = valueobjects.HalfMarathonLongDistancesMap[planLength][i]
	}

	return nil
}

/**
 * Easy run day is the first available day that is not long run day
 */
func setEasyRuns(
	weeks *[]model.ActivityWeek,
	rd valueobjects.RunningDays,
	longRunDay int,
	planLength int,
) error {
	for i, w := range *weeks {
		if i == len(*weeks)-1 {
			return nil
		}

		easyIndex, err := w.GetEasyRunDay(rd, longRunDay)
		if err != nil {
			return err
		}

		erd, err := w.GetDayByIndex(easyIndex)
		if err != nil {
			return err
		}

		erd.Distance = valueobjects.HalfMarathonLongDistancesMap[planLength][i]
	}

	return nil
}

func setThresholdRuns(weeks *[]model.ActivityWeek) {

}
