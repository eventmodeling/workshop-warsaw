package register

import (
	"crypto/sha256"
	"fmt"

	"github.com/pkg/errors"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

const UserRegisteredEventName = "UserRegistered"

type Register struct {
	Name     string
	Email    string
	Password string
}

type UserRegistered struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type RegisterHandler struct {
	Publisher    publisher
	EventsReader eventsReader
}

func (h RegisterHandler) Execute(cmd Register) error {
	hash := sha256.Sum256([]byte(cmd.Password))

	event := UserRegistered{
		Name:         cmd.Name,
		Email:        cmd.Email,
		PasswordHash: fmt.Sprintf("%x", hash),
	}

	// check prior registrations
	registered, err := h.allRegistrations()
	if err != nil {
		return err
	}
	for _, reg := range registered {
		if reg.Email == cmd.Email {
			return ErrUserAlreadyRegistered
		}
	}

	err = h.Publisher.Publish(UserRegisteredEventName, event)
	if err != nil {
		return err
	}

	return nil
}

func (h RegisterHandler) allRegistrations() ([]UserRegistered, error) {
	cur, err := h.EventsReader.ByName(UserRegisteredEventName)
	if errors.Cause(err) == ErrNoMatchesByName {
		return []UserRegistered{}, nil
	}

	ret := make([]UserRegistered, cur.Len())
	for i := 0; ; i++ {
		var event UserRegistered
		err = cur.Unmarshal(&event)
		if err != nil {
			return nil, err
		}
		ret[i] = event

		if !cur.Next() {
			break
		}
	}
	return ret, nil
}
