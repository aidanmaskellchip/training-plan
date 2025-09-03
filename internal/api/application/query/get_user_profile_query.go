package query

import (
	"log"
	"sync"
	model2 "training-plan/internal/api/domain/model"
	"training-plan/internal/api/domain/plan/entities"
	"training-plan/internal/api/domain/value_objects"
	"training-plan/internal/api/infrastructure/repository"
	response2 "training-plan/internal/api/transport/response"
)

func GetUserProfileQuery(id *string, repos *repository.Repositories) (res response2.GetUserProfileResponse, err error) {
	userID := valueobjects.NewUserID(*id)

	user, err := repos.UserRepository.FindByID(userID.ID)
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	actChan := make(chan model2.ActivityStats, 1)
	runProfChan := make(chan model2.RunningProfile, 1)
	favRunChan := make(chan entities.ActivityType, 1)

	go func() {
		defer wg.Done()

		longRun, err := repos.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		actChan <- longRun
	}()

	go func() {
		defer wg.Done()

		favRun, err := repos.GetMostCommonActivityType(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		favRunChan <- favRun
	}()

	go func() {
		defer wg.Done()

		runProf, err := repos.FindLatestUserProfile(userID.ID)
		if err != nil {
			log.Println(err)
			return
		}

		runProfChan <- runProf
	}()

	wg.Wait()

	res.Username = user.Username
	res.JoinedDate = user.CreatedAt.String()

	rp, err := response2.NewFindRunningProfileResponse(<-runProfChan)
	if err != nil {
		log.Println(err)

		return
	}
	res.LatestRunningProfile = rp

	res.LongestRun = <-actChan
	res.MostCommonActivityType = <-favRunChan

	return res, nil

}
