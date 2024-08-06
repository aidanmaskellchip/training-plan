package query

import (
	"errors"
	"log"
	"sync"
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
)

func GetUserStatsQuery(id *string, repos *repository.Repositories) (res []model.ActivityStats, err error) {
	userID := vo.NewUserID(*id)

	statsChan := make(chan model.ActivityStats)
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		act, err := repos.UserActivityRepository.GetFastestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	go func() {
		act, err := repos.UserActivityRepository.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	go func() {
		act, err := repos.UserActivityRepository.GetFastestCommunityActivity()
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	go func() {
		act, err := repos.UserActivityRepository.GetLongestCommunityActivity()
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	wg.Wait()

	if len(statsChan) != 4 {
		return res, errors.New("could not retrieve stats")
	}

	for i := 0; i < 4; i++ {
		res = append(res, <-statsChan)
	}

	return res, nil
}
