package events

import "encoding/json"

type UserCreatedEvent struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func (u UserCreatedEvent) ToBytes() []byte {
	ev, _ := json.Marshal(u)

	return ev
}
