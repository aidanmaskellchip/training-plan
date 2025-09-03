package factory

import (
	"fmt"
	"sync"
	model2 "training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/plan/entities"
)

func NewPlan(rp model2.RunningProfile) (p model2.Plan, err error) {
	p.UserID = rp.UserID
	p.Length = rp.PlanLength
	p.GoalDistance = rp.GoalDistance

	var weeks []model2.ActivityWeek
	for i := 0; i < rp.PlanLength; i++ {
		weeks = append(weeks, NewWeek())
	}

	rd, err := entities.RunningDaysFromJson(rp.RunningDays)
	if err != nil {
		return model2.Plan{}, err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	errChn := make(chan error, 2)

	go func() {
		defer wg.Done()

		if err := setLongRuns(&weeks, rp.LongRunDay, rp.PlanLength); err != nil {
			errChn <- err
		}
	}()

	go func() {
		defer wg.Done()

		if err := setEasyRuns(&weeks, rd, rp.LongRunDay, rp.PlanLength); err != nil {
			errChn <- err
		}
	}()

	wg.Wait()
	close(errChn)

	for e := range errChn {
		if e != nil {
			return model2.Plan{}, e
		}
	}

	// long runs distances
	// easy run distances
	// threshold

	return p, nil
}

func setLongRuns(weeks *[]model2.ActivityWeek, longRunDay int, planLength int) error {
	if _, ok := entities.HalfMarathonLongDistancesMap[planLength]; !ok {
		return fmt.Errorf("invalid plan length: %d", planLength)
	}
	for i, w := range *weeks {
		if i == len(*weeks)-1 {
			return nil
		}

		lrdPtr, err := w.GetDayByIndex(longRunDay)
		if err != nil {
			return err
		}

		if i >= len(entities.HalfMarathonLongDistancesMap[planLength]) {
			return fmt.Errorf("index out of bounds for long run distances: %d", i)
		}
		lrdPtr.Distance = entities.HalfMarathonLongDistancesMap[planLength][i]
	}

	return nil
}

/**
 * Easy run day is the first available day that is not long run day
 */
func setEasyRuns(
	weeks *[]model2.ActivityWeek,
	rd entities.RunningDays,
	longRunDay int,
	planLength int,
) error {
	if _, ok := entities.HalfMarathonEasyDistancesMap[planLength]; !ok {
		return fmt.Errorf("invalid plan length: %d", planLength)
	}
	for i, w := range *weeks {
		if i == len(*weeks)-1 {
			return nil
		}

		easyIndex, err := w.GetEasyRunDay(rd, longRunDay)
		if err != nil {
			return err
		}

		erdPtr, err := w.GetDayByIndex(easyIndex)
		if err != nil {
			return err
		}

		if i >= len(entities.HalfMarathonEasyDistancesMap[planLength]) {
			return fmt.Errorf("index out of bounds for easy run distances: %d", i)
		}
		erdPtr.Distance = entities.HalfMarathonEasyDistancesMap[planLength][i]
	}

	return nil
}

func setThresholdRuns(weeks *[]model2.ActivityWeek) {

}
