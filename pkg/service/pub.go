package service

import (
	"sth"
	"sth/repository"
)

type PubService struct {
	repo repository.Publisher
}

func NewPubService(repo repository.Publisher) *PubService {
	return &PubService{repo: repo}
}

func (p *PubService)ShowById(id, id_sub int)([]sth.Message, error) {
	return p.repo.ShowById(id, id_sub)
}

func (p *PubService)ShowAll(id int)([]sth.Message, error) {
	return p.repo.ShowAll(id)
}