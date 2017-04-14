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
	//Presenter	*Presenter //TODO: should be a PresenterInput so we can swap em out if we ever want to!
}

//TODO: make this a shared adapter: level object
//type Connection struct {
//	Request *http.Request
//	Writer PresenterOutput
//}
//
//func (c *Controller) Route() {
//	router.Post("/user", c.create)
//	router.Get("/user", c.index)
//	router.Get("/user/:id", c.show)
//}

func (c *Controller) Index(r *http.Request) {
	//c.Presenter.Connection = Connection{Writer: w, Request: r}
	c.Output.GetAllUsers()
}

func (c *Controller) Create(r *http.Request) {
	//c.Presenter.Connection = Connection{Writer: w, Request: r}

	r.ParseForm()
	firstname := r.PostFormValue("firstname")
	lastname := r.PostFormValue("lastname")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := domain.User{FirstName: firstname, LastName: lastname, Email: email, Password: password}

	c.Output.CreateUser(user)
}

func (c *Controller) Show(r *http.Request) {
	//c.Presenter.Connection = Connection{Writer: w, Request: r}
	userID, err := strconv.Atoi(c.ParamExtractor.Param(r, "id"))

	if err != nil {
		c.Output.Error(err)
		return
	}

	//call interactor to get user by id
	c.Output.GetUser(userID)
}


