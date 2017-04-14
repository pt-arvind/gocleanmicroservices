package view

import (
	"net/http"
	"encoding/json"
	"domain"
	"cmd/dbservice/adapter/viewport"
)

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

func (f JSONViewFactory) NewJSONViewport() viewport.Viewport {
	return &JSONView{}
}

type JSONView struct {

}

func render(w http.ResponseWriter, users []UserJSONModel, errJSON *ErrorJSONModel) {
	jsonResponse := JSONResponse{}
	jsonResponse.Error = errJSON
	jsonResponse.Users = users

	jsonBytes, err := json.Marshal(jsonResponse)

	if err != nil {
		panic("should never happen: " + err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

/* viewport conformance */
func (v *JSONView) Render(w http.ResponseWriter, users []domain.User, err error) {
	jsonModels := []UserJSONModel{}

	if err != nil {
		errorJSON := (&ErrorJSONModel{}).from(err)
		render(w, jsonModels, &errorJSON)
		return
	}


	for _,user := range users {
		jsonModel := (&UserJSONModel{}).from(user)
		jsonModels = append(jsonModels, jsonModel)
	}

	render(w, jsonModels, nil)
}