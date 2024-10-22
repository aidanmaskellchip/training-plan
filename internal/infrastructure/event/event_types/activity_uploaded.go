package eventtypes

const ActivityUploadedEventTypeName = "activity_uploaded_event"

type ActivityUploadedEvent struct {
	UserID string `json:"id"`
}
