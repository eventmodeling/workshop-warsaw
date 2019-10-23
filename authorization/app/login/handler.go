package login

type Login struct {
	Email string
	Password string
}

type LoginHandler struct{}

func (h LoginHandler) Execute(cmd Login) error {
	return nil
}
