package dbservice

import "domain"

//FIXME: these are identical/shared between dbservice and webapp
func (u *UserJSONModel) toDomUser() domain.User {

	user := domain.User{}
	user.ID = u.ID
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.Password = u.Password

	return user
}


type UserJSONModel struct {
	ID        int		`json:"id"`
	FirstName string	`json:"firstname"`
	LastName  string	`json:"lastname"`
	Email     string	`json:"email"`
	Password  string	`json:"password"`
}
func (m *UserJSONModel) from(user domain.User) UserJSONModel {
	m.ID = user.ID
	m.FirstName = user.FirstName
	m.LastName = user.LastName
	m.Email = user.Email
	m.Password = user.Password

	return *m
}

type ErrorJSONModel struct {
	Message string	`json:"error_message"`
}
func (e *ErrorJSONModel) from(err error) ErrorJSONModel {
	e.Message = err.Error()
	return *e
}

type JSONResponse struct {
	Users []UserJSONModel	`json:"users"`
	Error *ErrorJSONModel	`json:"error,omitempty"`
}

type JSONViewFactory struct {

}

type JSONView struct {

}
