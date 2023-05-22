package service

import (
	"sth"
	"sth/pkg/repository"
)

type SubService struct {
	repo repository.Subscriber
}

func NewSubService(repo repository.Subscriber) *SubService {
	return &SubService{repo: repo}
}

func (p *SubService)ShowById(id, id_sub int)([]sth.Message, error) {
	return p.repo.ShowById(id, id_sub)
}

func (p *SubService)ShowAll(id int)([]sth.Message, error) {
	return p.repo.ShowAll(id)
}