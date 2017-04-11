package register

import (
	"net/http"
	"github.com/pt-arvind/gocleanarchitecture/logic"
)

// RegisterHandler represents the services required for this controller.
//type RegisterHandler struct {
//	UserService domain.UserInteractor
//	ViewService domain.ViewCase
//}

//type Controller struct {
//	//UserService domain.UserInteractor
//	//ViewService domain.ViewCase
//	Output  	InteractorInput
//}

//TODO: not the best spot to put this
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
		//controller.Output.RequestStore(conn)
		controller.createUser(writer, request)
	} else {
		//controller.Output.RequestIndex(conn)
		controller.index(writer, request)
	}
}

func (controller *Controller) index(writer http.ResponseWriter, request *http.Request) {
	controller.Output.Index()
}

func (controller *Controller) createUser(writer http.ResponseWriter, request *http.Request) {
	// Don't continue if required fields are missing.
	// validation
	//for _, v := range []string{"firstname", "lastname", "email", "password"} {
	//	if len(request.FormValue(v)) == 0 {
	//		//call presenter present400
	//		interactor.Output.Present400(conn)
	//		return
	//	}
	//}

	// Build the user from the form values.
	//u := new(domain.User)
	firstname := request.FormValue("firstname")
	lastname := request.FormValue("lastname")
	email := request.FormValue("email")
	password := request.FormValue("password")

	//fmt.Println(firstname)
	//fmt.Println(lastname)
	//fmt.Println(email)
	//fmt.Println(password)

	// Add the user to the database.
	//err := interactor.UserService.CreateUser(u)

	controller.Output.CreateUser(firstname,lastname,email,password)

	//if err != nil {
	//	//call presenter present500
	//	interactor.Output.Present500(conn, err)
	//} else {
	//	//call presenter presentSuccess
	//	interactor.Output.PresentSuccessfulUserCreation(conn)
	//}
}