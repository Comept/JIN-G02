package service

import (
	"sth"
	"sth/pkg/repository"
)

//const salt = "sofdsoifjeoifjoewijffdg"
type Authorization interface {
	CreateUser(user sth.Users) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Message interface {
	Create(message sth.Message) (int, error)
	//GetUser(username, password string) (string, error)
}

type Publisher interface {
	ShowById(id, id_sub int) ([]sth.Message, error)
	ShowAll(id int)([]sth.Message, error)
}
type Subscriber interface {
	ShowById(id, id_pub int) ([]sth.Message, error)
	ShowAll(id int)([]sth.Message, error)
}

type Service struct {
	Authorization
	Message
	Publisher
	Subscriber
}

func NewService (repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Message: NewMessageService(repo.Message),
		Publisher: NewPubService(repo.Publisher),
		Subscriber: NewSubService(repo.Subscriber),
	}
}

