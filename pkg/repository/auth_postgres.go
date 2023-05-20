package repository

import (
	"fmt"
	"sth"

	"github.com/go-pg/pg/v10"
)

type AuthPostgres struct {
	db *pg.DB
}

func NewAuthPostgres(db *pg.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (s *AuthPostgres) CreateUser(user sth.Users) (int, error) {
	


	_, err := s.db.Model(&user).Insert()
	if err != nil {
		//handler.newErrorResponse(err)
		fmt.Println("sfsdfsfdsd")
		return 0, err


	}
	fmt.Print(user.Id)
	return user.Id, nil
}

func (c *AuthPostgres) GetUser(username, password string) (sth.Users, error){
	var user sth.Users
	fmt.Println(username)
	fmt.Println("username")
	err := c.db.Model(&user).Where("? = ?", pg.Ident("username"), username).Where("? = ?", pg.Ident("password"), password).Select()
	if err != nil {
		panic(err)
	}
	return user, err
}
