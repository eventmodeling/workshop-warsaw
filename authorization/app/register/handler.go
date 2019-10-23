package register

type Register struct {
	Email string
	Password string
}

type RegisterHandler struct{}

func (h RegisterHandler) Execute(cmd Register) error {
	return nil
}
