package query

import (
	"errors"
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

	actChan := make(chan model.ActivityStats)
	runProfChan := make(chan model.RunningProfile)

	go func() {
		defer wg.Done()

		longRun, err := repos.UserActivityRepository.GetLongestUserActivity(userID.ID)
		if err != nil {
			log.Println(err)
			// do i need this ? or will the defer one cover it
			wg.Done()
			return
		}

		actChan <- longRun
	}()

	go func() {
		defer wg.Done()

		favRun, err := repos.UserActivityRepository.GetMostCommonActivityType(userID.ID)
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		actChan <- favRun
	}()

	go func() {
		defer wg.Done()

		runProf, err := repos.RunningProfileRepository.FindLatestUserProfile(userID.ID)
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}

		runProfChan <- runProf
	}()

	wg.Wait()

	if len(actChan) != 2 {
		return res, errors.New("could not retrieve profile stats")
	}

	res.Username = user.Username
	res.JoinedDate = user.CreatedAt.String()

	rp, err := response.NewFindRunningProfileResponse(<-runProfChan)
	if err != nil {
		log.Println(err)

		return
	}
	res.LatestRunningProfile = rp

	if len(actChan) != 2 {
		log.Println(err)

		return res, errors.New("could not retrieve user activity stats")
	}

	act := <-actChan
	if act.Title == vo.STATS_TYPE_USER_MOST_COMMON_ACTIVITY {
		res.MostCommonActivityType = act.Type
		res.LongestRun = <-actChan
	} else {
		res.LongestRun = act
		favAct := <-actChan
		res.MostCommonActivityType = favAct.Type
	}

	// username
	// joined date
	// longest run
	// fav type

	return res, nil

}
