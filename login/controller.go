package login

import (
	"net/http"
)

type Controller struct {
	Output  	InteractorInput
}

//TODO: not the best spot to put this
type Connection struct {
	Request *http.Request
	Writer PresenterOutput
}

type ViewModel struct {

}


// Index displays the logon screen.
func (controller *Controller) Index(w http.ResponseWriter, r *http.Request) {
	connection := Connection{Request: r, Writer: w}
	// Handle 404.
	if r.URL.Path != "/" { // FIXME: will typically be handled by a router, so it's OK for this logic to be in here for now
		controller.Output.Request404(connection)
	} else if r.Method == "POST" { 	// FIXME: will typically be handled by a router, so it's OK for this logic to be in here for now
		// call store on interactor
		controller.Output.RequestStore(connection)
	} else {
		controller.Output.RequestIndex(connection)
	}
}

// Store handles the submission of the login information.
//func (h *Controller) Store(w http.ResponseWriter, r *http.Request) {
//}
