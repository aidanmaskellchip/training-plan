package query

import (
	"errors"
	"log"
	"sync"
	"training-plan/internal/api/domain/model"
	vo "training-plan/internal/api/domain/value_objects"
	"training-plan/internal/api/infrastructure/repository"
)

func GetUserStatsQuery(id *string, repos *repository.Repositories) (res []model.ActivityStats, err error) {
	userID := vo.NewUserID(*id)

	statsChan := make(chan model.ActivityStats, 4)
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		act, err := repos.GetFastestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	go func() {
		act, err := repos.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	go func() {
		act, err := repos.GetFastestCommunityActivity()
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		statsChan <- act
		wg.Done()
	}()

	go func() {
		act, err := repos.GetLongestCommunityActivity()
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
