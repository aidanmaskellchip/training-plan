package planfactory

import (
	"sync"
	weekfactory "training-plan/internal/domain/factory/week_factory"
	"training-plan/internal/domain/model"
	valueobjects "training-plan/internal/domain/value_objects"
)

func NewPlan(rp model.RunningProfile) (p model.Plan, err error) {
	var weeks []model.ActivityWeek
	for i := 0; i < rp.PlanLength; i++ {
		weeks = append(weeks, weekfactory.NewWeek())
	}

	//TODO: incorporate new running days function then pass structure into these set runs methods
	//rd, err := valueobjects.RunningDaysFromJson(rp.RunningDays)

	wg := sync.WaitGroup{}
	wg.Add(1)
	errChn := make(chan error)

	go func() {
		err = setLongRuns(&weeks, rp.LongRunDay, rp.PlanLength)
		if err != nil {
			errChn <- err
		}

		wg.Done()
	}()

	go func() {
		rd, err := valueobjects.RunningDaysFromJson(rp.RunningDays)
		if err != nil {
			errChn <- err
			wg.Done()

			return
		}

		err = setEasyRuns(&weeks, rd, rp.LongRunDay, rp.PlanLength)
		if err != nil {
			errChn <- err
		}

		wg.Done()
	}()

	wg.Wait()

	// long runs distances
	// easy run distances
	// threshold

	return model.Plan{}, nil
}

func setLongRuns(weeks *[]model.ActivityWeek, longRunDay int, planLength int) error {
	for i, w := range *weeks {
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

func setThresholdRuns(weeks *[]model.ActivityWeek)
