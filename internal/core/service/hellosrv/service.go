package hellosrv

import "github.com/victoorraphael/poc-hex-shelf/internal/core/ports"

type service struct {
	repo ports.HelloRepository
}

func New(repo ports.HelloRepository) *service {
	return &service{repo: repo}
}

func (s *service) Get() string {
	return s.repo.Get()
}
