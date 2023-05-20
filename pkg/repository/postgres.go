package repository

import (
	"github.com/go-pg/pg/v10"
)


type Config struct {
	Addr  string
	User  string
	Password  string
	Database  string
}

func NewPostgresDB(cfg Config) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Addr,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})
	return db 
}