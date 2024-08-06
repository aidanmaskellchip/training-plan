package query

import (
	"log"
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
)

func GetUserStatsQuery(id *string, repos *repository.Repositories) (res []model.ActivityStats, err error) {
	userID := vo.NewUserID(*id)

	statsChan := make(chan model.ActivityStats)

	go func() {
		act, err := repos.UserActivityRepository.GetFastestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	go func() {
		act, err := repos.UserActivityRepository.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	go func() {
		act, err := repos.UserActivityRepository.GetFastestCommunityActivity()
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	go func() {
		act, err := repos.UserActivityRepository.GetLongestCommunityActivity()
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	for i := 0; i < 4; i++ {
		res = append(res, <-statsChan)
	}

	return res, nil
}
