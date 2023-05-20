package repository

import (
	"fmt"
	"sth"

	"github.com/go-pg/pg/v10"
)

type MessagePostgres struct {
	db *pg.DB
}

func NewMessagePostgres(db *pg.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

func (r *MessagePostgres) Create(message sth.Message) (int, error){
	fmt.Println(message.Sub)
	fmt.Println("asdasd")
	fmt.Println(message.Message)
	message.Message = "asdasd"
	fmt.Println(message.Message)
	fmt.Println(message.Pub)
	_, err := r.db.Model(&message).Insert()
	if err != nil {
		panic(err)
	}
	fmt.Print(message.Id)
	return message.Id, nil
}