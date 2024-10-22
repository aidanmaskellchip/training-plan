package event

var ActivityUploadedCh chan Event

type Channels struct {
	ActivityUploadedCh chan Event
}

func NewChannels() *Channels {
	return &Channels{
		ActivityUploadedCh: ActivityUploadedCh,
	}
}
