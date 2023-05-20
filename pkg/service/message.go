package service

import (
	"sth"
	"sth/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessageService(repo repository.Message) *MessageService{
	return &MessageService{repo:repo}
}
func (c *MessageService) Create( message sth.Message) (int, error){
	return c.repo.Create(message)
}
