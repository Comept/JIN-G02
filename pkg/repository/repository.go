package repository

import (
	"sth"

	"github.com/go-pg/pg/v10"
)

type Authorization interface {
	CreateUser(user sth.Users) (int, error)
	GetUser(username, password string) (sth.Users, error)
}

type Message interface {
	Create(message sth.Message) (int, error)
}

type Publisher interface {
	ShowById(id, id_sub int) ([]sth.Message, error)
	ShowAll(id int)([]sth.Message, error)
}
type Subscriber interface {
	ShowById(id, id_pub int) ([]sth.Message, error)
	ShowAll(id int)([]sth.Message, error)
}

type Repository struct {
	Authorization
	Message
	Publisher
	Subscriber
}

func NewRepository(db *pg.DB) *Repository{
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Message: NewMessagePostgres(db),
		Publisher: NewPubPostgres(db),
		Subscriber: NewSubPostgres(db),
	}

}

