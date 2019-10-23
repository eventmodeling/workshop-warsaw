package register

import (
	"crypto/sha256"
	"fmt"
)

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

type publisher interface {
	Publish(eventName string, event interface{}) error
}

type RegisterHandler struct {
	Publisher publisher
}

func (h RegisterHandler) Execute(cmd Register) error {
	hash := sha256.Sum256([]byte(cmd.Password))

	event := UserRegistered{
		Name:         cmd.Name,
		Email:        cmd.Email,
		PasswordHash: fmt.Sprintf("%x", hash),
	}

	err := h.Publisher.Publish("UserRegistered", event)
	if err != nil {
		return err
	}

	return nil
}
