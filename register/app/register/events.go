package register

import "github.com/pkg/errors"

var (
	ErrNoEventsByName = errors.New("no events found by name")
)

const UserRegisteredEventName = "UserRegistered"

type UserRegistered struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type publisher interface {
	Publish(eventName string, event interface{}) error
}

type eventsReader interface {
	ByName(eventName string) (EventsCursor, error)
}

type EventsCursor interface {
	// Len returns the length of the cursor.
	Len() int
	// Next bumps the iterator and returns false if there is no more events
	Next() bool
	// Unmarshal unmarshals the current event on the given structure.
	Unmarshal(val interface{}) error
}
