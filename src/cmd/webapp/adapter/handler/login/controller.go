package login

import (
	"net/http"
	"cmd/webapp/logic"
	"errors"
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

// Index displays the logon screen.
func (controller *Controller) Route(writer http.ResponseWriter, request *http.Request) {

	// set connection on presenter :(
	// TODO: perhaps a router could handle doing this part for us and remove the reference between controller and presenter?
	controller.Presenter.Connection = Connection{Writer: writer, Request: request}

	if request.URL.Path != "/" {
		controller.error404(writer,request)

	} else if request.Method == "POST" {
		controller.authenticate(writer,request)
	} else {
		controller.index(writer,request)
	}
}

func (controller *Controller) error404(writer http.ResponseWriter, request *http.Request) {
	controller.Output.Error(errors.New("404"))
}

func (controller *Controller) index(writer http.ResponseWriter, request *http.Request) {
	controller.Output.Index()
}

// Store handles the submission of the login information.
func (controller *Controller) authenticate(writer http.ResponseWriter, request *http.Request) {

	for _, v := range []string{"email", "password"} {
		if len(request.FormValue(v)) == 0 {
			controller.Output.Error(errors.New("fill out all form fields prior to submitting!"))
			return
		}
	}

	email := request.FormValue("email")
	password := request.FormValue("password")

	controller.Output.Authenticate(email, password)
}
