package user

import (
	"net/http"
	"github.com/blue-jay/core/router" //FIXME: use an interface to inject it
	"cmd/dbservice/logic"
	"domain"
	"strconv"
)

type Controller struct {
	Output  	logic.UserInteractorInput
	Presenter	*Presenter //TODO: should be a PresenterInput so we can swap em out if we ever want to!
}

//TODO: make this a shared adapter: level object
type Connection struct {
	Request *http.Request
	Writer PresenterOutput
}

func (c *Controller) Route() {
	router.Post("/user", c.create)
	router.Get("/user", c.index)
	router.Get("/user/:id", c.show)
}

func (c *Controller) index(w http.ResponseWriter, r *http.Request) {
	c.Presenter.Connection = Connection{Writer: w, Request: r}
	c.Output.GetAllUsers()
}

func (c *Controller) create(w http.ResponseWriter, r *http.Request) {
	c.Presenter.Connection = Connection{Writer: w, Request: r}

	r.ParseForm()
	firstname := r.PostFormValue("firstname")
	lastname := r.PostFormValue("lastname")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	user := domain.User{FirstName: firstname, LastName: lastname, Email: email, Password: password}

	c.Output.CreateUser(user)
}

func (c *Controller) show(w http.ResponseWriter, r *http.Request) {
	c.Presenter.Connection = Connection{Writer: w, Request: r}
	userID, err := strconv.Atoi(router.Param(r, "id"))

	if err != nil {
		c.Output.Error(err)
		return
	}

	//call interactor to get user by id
	c.Output.GetUser(userID)
}


