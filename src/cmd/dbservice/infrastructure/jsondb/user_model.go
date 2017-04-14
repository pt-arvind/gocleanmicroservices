package jsondb

import (
	"domain"
)

type User struct {
	ID        int		`json:"id"`
	FirstName string	`json:"firstname"`
	LastName  string	`json:"lastname"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}

func (u *User) from (user domain.User) User {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Password = user.Password

	return *u
}

func (u *User) toDomUser() domain.User {

	user := domain.User{}
	user.ID = u.ID
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.Password = u.Password

	return user
}


func DBtoDomain(users []User) []domain.User {

	domUsers := []domain.User{}
	for _, user := range users {
		domUser := user.toDomUser()
		domUsers = append(domUsers, domUser)
	}

	return domUsers
}
//
//func DomainToDB(users []domain.User) []User{
//
//}
//
