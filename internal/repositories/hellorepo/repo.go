package hellorepo

type hellorepo struct{}

func NewRepo() *hellorepo {
	return &hellorepo{}
}

func (h *hellorepo) Get() string {
	return "hello"
}
