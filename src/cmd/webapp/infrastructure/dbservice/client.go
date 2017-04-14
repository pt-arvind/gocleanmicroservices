package dbservice

import (
	"domain"
	"net/http"
	"net/url"
	"strings"
	"encoding/json"
)

//type ServiceInput interface {
//	Records(output func([]domain.User))
//	AddRecord(user domain.User, output func(user *domain.User, err error))
//}

/* ServiceInput conformance */

type Client struct {
	Client http.Client
}

func (c Client) Records(output func(users []domain.User, err error)) {
	req, err := http.NewRequest("GET", "http://localhost:8081/user", nil)

	if err != nil {
		output(nil, err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// FIXME: do it async
	resp, err := c.Client.Do(req)


	if err != nil {
		output(nil, err)
		return
	}
	defer resp.Body.Close()

	jsonResponse := JSONResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		output(nil, err)
		return
	}

	users := []domain.User{}

	for _, userJSON := range jsonResponse.Users {
		user := userJSON.toDomUser()
		users = append(users, user)
	}

	output(users, nil)


}

func (c Client) AddRecord(user domain.User, output func(user *domain.User, err error)) {
	form := url.Values{}
	form.Add("firstname", user.FirstName)
	form.Add("lastname", user.LastName)
	form.Add("email", user.Email)
	form.Add("password", user.Password)

	//FIXME: make this stuff injectable/configurable
	req, err := http.NewRequest("POST", "http://localhost:8081/user", strings.NewReader(form.Encode()))
	req.PostForm = form

	if err != nil {
		output(nil, err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// FIXME: do it async
	resp, err := c.Client.Do(req)

	if err != nil {
		output(nil, err)
		return
	}
	defer resp.Body.Close()

	jsonResponse := JSONResponse{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		output(nil, err)
		return
	}

	createdUser := jsonResponse.Users[0].toDomUser()
	output(&createdUser, nil)

}

//FIXME: find me a better home
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
