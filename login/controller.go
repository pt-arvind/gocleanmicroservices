package login

import (
	"net/http"
	"github.com/pt-arvind/gocleanarchitecture/logic"
	"errors"
)

type Controller struct {
	Output  	logic.UserInteractorInput
	Presenter	*Presenter
}

//TODO: not the best spot to put this
type Connection struct {
	Request *http.Request
	Writer PresenterOutput
}


// Index displays the logon screen.
func (controller *Controller) Route(writer http.ResponseWriter, request *http.Request) {

	// set connection on presenter :(
	// TODO: would love a better solution than this...
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
	// call store on interactor
	//controller.Output.RequestStore(connection)


	// validation of this level should honestly take place in javascript or prior to even getting here!
	//for _, v := range []string{"email", "password"} {
	//	if len(conn.Request.FormValue(v)) == 0 {
	//		interactor.Output.Present400(conn)
	//		return
	//	}
	//}


	//user := new(domain.User)


	email := request.FormValue("email")
	password := request.FormValue("password")

	//err := interactor.UserInteractor.Authenticate(user)

	//if err != nil {
	//	interactor.Output.Present401(conn) //realistically would probably pass that error along
	//} else {
	//	interactor.Output.PresentSuccessfulLogin(conn) // realistically, you'd want to have something here that would pass along the user you just made
	//}

	controller.Output.Authenticate(email, password)

}
