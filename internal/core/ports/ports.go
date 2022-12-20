package ports

type HelloRepository interface {
	Get() string
}

type HelloService interface {
	Get() string
}
