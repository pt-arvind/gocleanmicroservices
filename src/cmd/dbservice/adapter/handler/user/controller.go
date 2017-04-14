package user

import (
	"net/http"
	"cmd/dbservice/logic"
	"domain"
	"strconv"
)

type ParamExtractor interface {
	Param(r *http.Request, name string) string
}

type Controller struct {
	Output  	logic.UserInteractorInput
	ParamExtractor	ParamExtractor
}

func (c *Controller) Index(r *http.Request) {
	c.Output.GetAllUsers()
}

func (c *Controller) Create(r *http.Request) {
	r.ParseForm()
	firstname := r.PostFormValue("firstname")
	lastname := r.PostFormValue("lastname")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := domain.User{FirstName: firstname, LastName: lastname, Email: email, Password: password}

	c.Output.CreateUser(user)
}

func (c *Controller) Show(r *http.Request) {
	userID, err := strconv.Atoi(c.ParamExtractor.Param(r, "id"))

	if err != nil {
		c.Output.Error(err)
		return
	}

	c.Output.GetUser(userID)
}


