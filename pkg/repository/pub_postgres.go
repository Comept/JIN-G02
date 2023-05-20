package repository

import (
	"fmt"
	"sth"

	"github.com/go-pg/pg/v10"
)

type PubPostgres struct {
	db *pg.DB
}

func NewPubPostgres(db *pg.DB) *PubPostgres {
	return &PubPostgres{db: db}
}

func (p *PubPostgres)ShowById(id, id_sub int)([]sth.Message, error) {
	var messages []sth.Message
	err := p.db.Model(&messages).Where("? = ?", pg.Ident("pub"), id).Where("? = ?", pg.Ident("sub"), id_sub).Select()
	if err != nil {
		return nil, err
	}
	fmt.Println(messages)
	return messages, err
}

func (p *PubPostgres)ShowAll(id int)([]sth.Message, error) {
	var messages []sth.Message
	err := p.db.Model(&messages).Where("? = ?", pg.Ident("pub"), id).Select()
	if err != nil {
		return nil, err
	}
	fmt.Println(messages)
	return messages, err
}