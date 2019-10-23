package feedback

import "errors"

type Level struct {
	level string
}

func (l Level) String() string {
	return l.level
}

var (
	LevelSad     = Level{"sad"}
	LevelNeutral = Level{"neutral"}
	LevelHappy   = Level{"happy"}
)

func NewLevel(name string) (Level, error) {
	switch name {
	case "sad":
		return LevelSad, nil
	case "neutral":
		return LevelNeutral, nil
	case "happy":
		return LevelHappy, nil
	}

	return Level{}, errors.New("no such level: " + name)
}

type Feedback struct {
	PurchaseID  string
	Cleanliness Level
	Experience  Level
	Queue       Level
	EasyToFind  Level
	Comments    string
}

type FeedbackHandler struct {
	Publisher publisher
}

func (h FeedbackHandler) Execute(cmd Feedback) error {
	event := FeedbackSubmitted{
		PurchaseID:  cmd.PurchaseID,
		Cleanliness: cmd.Cleanliness.String(),
		Experience:  cmd.Experience.String(),
		Queue:       cmd.Queue.String(),
		EasyToFind:  cmd.EasyToFind.String(),
		Comments:    cmd.Comments,
	}

	err := h.Publisher.Publish(FeedbackSubmittedEventName, event)
	if err != nil {
		return err
	}

	return nil
}
