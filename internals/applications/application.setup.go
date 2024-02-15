package applications

type Storage struct{}

func New() *Storage {
	return &Storage{}
}

type HelloWorldHandler1Request struct {
	Email string `json:"email"`
}
