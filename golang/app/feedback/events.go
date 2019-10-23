package feedback

type FeedbackSubmitted struct {
	Cleanliness string `json:"cleanliness"`
	Experience  string `json:"experience"`
	Queue       string `json:"queue"`
	EasyToFind  string `json:"easy_to_find"`
	Comments    string `json:"comments"`
}

const FeedbackSubmittedEventName = "FeedbackSubmitted"

type publisher interface {
	Publish(eventName string, event interface{}) error
}
