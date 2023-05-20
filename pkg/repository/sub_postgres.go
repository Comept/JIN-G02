package repository

import (
	"fmt"
	"sth"

	"github.com/go-pg/pg/v10"
)

type SubPostgres struct {
	db *pg.DB
}

func NewSubPostgres(db *pg.DB) *SubPostgres {
	return &SubPostgres{db: db}
}

func (p *SubPostgres) ShowById(id, id_pub int) ([]sth.Message, error) {
	var messages []sth.Message
	err := p.db.Model(&messages).Where("? = ?", pg.Ident("sub"), id).Where("? = ?", pg.Ident("pub"), id_pub).Select()
	if err != nil {
		return nil, err
	}
	fmt.Println(messages)
	return messages, err
}

func (p *SubPostgres) ShowAll(id int) ([]sth.Message, error) {
	var messages []sth.Message
	err := p.db.Model(&messages).Where("? = ?", pg.Ident("sub"), id).Select()
	if err != nil {
		return nil, err
	}
	fmt.Println(messages)
	return messages, err
}