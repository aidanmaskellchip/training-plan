package model

import (
	"errors"
	valueobjects2 "training-plan/internal/api/domain/value_objects"
)

type ActivityWeek struct {
	Mon valueobjects2.Activity
	Tue valueobjects2.Activity
	Wed valueobjects2.Activity
	Thu valueobjects2.Activity
	Fri valueobjects2.Activity
	Sat valueobjects2.Activity
	Sun valueobjects2.Activity
}

func (aw *ActivityWeek) GetDayByIndex(i int) (*valueobjects2.Activity, error) {
	switch i {
	case 0:
		return &aw.Mon, nil
	case 1:
		return &aw.Tue, nil
	case 2:
		return &aw.Wed, nil
	case 3:
		return &aw.Thu, nil
	case 4:
		return &aw.Fri, nil
	case 5:
		return &aw.Sat, nil
	case 6:
		return &aw.Sun, nil
	default:
		return nil, errors.New("invalid activity index")
	}
}

func (aw ActivityWeek) GetEasyRunDay(days valueobjects2.RunningDays, longRunDay int) (int, error) {
	for i, v := range days.Days {
		if v == 0 || i == longRunDay {
			continue
		}

		return i, nil
	}

	return 0, errors.New("no free activity days")
}
