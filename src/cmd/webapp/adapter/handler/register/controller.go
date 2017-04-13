package register

import (
	"net/http"
	"cloudtamer/portal/cmd/webapp/logic"
	"github.com/pkg/errors"
)

//TODO: make this a shared adapter: level object
type Connection struct {
	Request *http.Request
	Writer PresenterOutput
}

type Controller struct {
	Output  	logic.UserInteractorInput
	Presenter	*Presenter
}


// Index displays the register screen.
func (controller *Controller) Route(writer http.ResponseWriter, request *http.Request) {
	controller.Presenter.Connection = Connection{Writer: writer, Request: request}

	if request.Method == "POST" { //will ideally be handled by a router
		//tests(tobemoved).Output.RequestStore(conn)
		controller.createUser(writer, request)
	} else {
		//tests(tobemoved).Output.RequestIndex(conn)
		controller.index(writer, request)
	}
}

func (controller *Controller) index(writer http.ResponseWriter, request *http.Request) {
	controller.Output.Index()
}

func (controller *Controller) createUser(writer http.ResponseWriter, request *http.Request) {

	//Don't continue if required fields are missing.
	 //validation
	for _, v := range []string{"firstname", "lastname", "email", "password"} {
		if len(request.FormValue(v)) == 0 {
			controller.Output.Error(errors.New("invalid input! you must fill out all the form fields!"))
			return
		}
	}

	firstname := request.FormValue("firstname")
	lastname := request.FormValue("lastname")
	email := request.FormValue("email")
	password := request.FormValue("password")

	controller.Output.CreateUser(firstname,lastname,email,password)
}