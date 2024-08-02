package query

import (
	"errors"
	"log"
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/response"
)

func GetUserStatsQuery(id *string, repos *repository.Repositories) (res *response.GetUserStatsResponse, err error) {
	userID := vo.NewUserID(*id)

	uFastChan := getUserFastestActivity(userID, repos)
	uFastAct, ok := <-uFastChan

	if !ok {
		err = errors.New("get user fastest activity failed")
		return
	}

	cFastChan := getCommunityFastestActivity(repos)
	cFastAct, ok := <-cFastChan

	if !ok {
		err = errors.New("get community fastest activity failed")
		return
	}

	uLongChan := getUserLongestActivity(userID, repos)
	uLongAct, ok := <-uLongChan

	if !ok {
		err = errors.New("get user longest activity failed")
		return
	}

	cLongChan := getCommunityLongestActivity(repos)
	cLongAct, ok := <-cLongChan

	if !ok {
		err = errors.New("get community longest activity failed")
		return
	}

	return &response.GetUserStatsResponse{
		UserFastestRun:      uFastAct,
		UserLongestRun:      uLongAct,
		CommunityFastestRun: cFastAct,
		CommunityLongestRun: cLongAct,
	}, nil
}

func getUserFastestActivity(userID vo.UserID, repos *repository.Repositories) <-chan model.ActivityStats {
	statsChan := make(chan model.ActivityStats)

	go func() {
		act, err := repos.UserActivityRepository.GetFastestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	return statsChan
}

func getCommunityFastestActivity(repos *repository.Repositories) <-chan model.ActivityStats {
	statsChan := make(chan model.ActivityStats)

	go func() {
		act, err := repos.UserActivityRepository.GetFastestCommunityActivity()
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	return statsChan
}

func getUserLongestActivity(userID vo.UserID, repos *repository.Repositories) <-chan model.ActivityStats {
	statsChan := make(chan model.ActivityStats)

	go func() {
		act, err := repos.UserActivityRepository.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	return statsChan
}

func getCommunityLongestActivity(repos *repository.Repositories) <-chan model.ActivityStats {
	statsChan := make(chan model.ActivityStats)

	go func() {
		act, err := repos.UserActivityRepository.GetLongestCommunityActivity()
		if err != nil {
			log.Println(err)
			return
		}

		statsChan <- act
	}()

	return statsChan
}
