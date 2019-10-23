package register

import "github.com/pkg/errors"

var (
	ErrCursorEnd       = errors.New("end of cursor")
	ErrNoMatchesByName = errors.New("no matches by name")
)

type publisher interface {
	Publish(eventName string, event interface{}) error
}

type eventsReader interface {
	ByName(eventName string) (EventsCursor, error)
}

type EventsCursor interface {
	Len() int
	Next() bool
	Unmarshal(val interface{}) error
}
