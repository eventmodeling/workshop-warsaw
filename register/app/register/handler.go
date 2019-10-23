package register

type Register struct {
	Name     string
	Email    string
	Password string
}

type UserRegistered struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type publisher interface {
	Publish(event interface{}) error
}

type RegisterHandler struct {
	Publisher publisher
}

func (h RegisterHandler) Execute(cmd Register) error {
	event := UserRegistered{
		Name:  cmd.Name,
		Email: cmd.Email,
	}

	err := h.Publisher.Publish(event)
	if err != nil {
		return err
	}

	return nil
}
