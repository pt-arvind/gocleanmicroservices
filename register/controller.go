package register

import (
	"net/http"
)

// RegisterHandler represents the services required for this controller.
//type RegisterHandler struct {
//	UserService domain.UserCase
//	ViewService domain.ViewCase
//}

type Controller struct {
	//UserService domain.UserCase
	//ViewService domain.ViewCase
	Output  	InteractorInput
}

//TODO: not the best spot to put this
type Connection struct {
	Request *http.Request
	Writer PresenterOutput
}


// Index displays the register screen.
func (controller *Controller) Index(w http.ResponseWriter, r *http.Request) {
	conn := Connection{Request: r, Writer: w}

	if r.Method == "POST" { //will ideally be handled by a router
		controller.Output.RequestStore(conn)
		return
	} else {
		controller.Output.RequestIndex(conn)
	}
}