package query

import (
	"log"
	"sync"
	"training-plan/internal/domain/model"
	vo "training-plan/internal/domain/value_objects"
	"training-plan/internal/infrastructure/repository"
	"training-plan/internal/transport/response"
)

func GetUserProfileQuery(id *string, repos *repository.Repositories) (res response.GetUserProfileResponse, err error) {
	userID := vo.NewUserID(*id)

	user, err := repos.UserRepository.FindByID(userID.ID)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	actChan := make(chan model.ActivityStats, 1)
	runProfChan := make(chan model.RunningProfile, 1)
	favRunChan := make(chan vo.ActivityType, 1)

	go func() {
		defer wg.Done()

		longRun, err := repos.UserActivityRepository.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		actChan <- longRun
	}()

	go func() {
		defer wg.Done()

		favRun, err := repos.UserActivityRepository.GetMostCommonActivityType(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		favRunChan <- favRun
	}()

	go func() {
		defer wg.Done()

		runProf, err := repos.RunningProfileRepository.FindLatestUserProfile(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		runProfChan <- runProf
	}()

	wg.Wait()

	res.Username = user.Username
	res.JoinedDate = user.CreatedAt.String()

	rp, err := response.NewFindRunningProfileResponse(<-runProfChan)
	if err != nil {
		log.Println(err)

		return
	}
	res.LatestRunningProfile = rp

	res.LongestRun = <-actChan
	res.MostCommonActivityType = <-favRunChan

	return res, nil

}
