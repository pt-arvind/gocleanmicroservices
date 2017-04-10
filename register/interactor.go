package register

import (
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

type InteractorInput interface {
	RequestIndex(conn Connection)
	RequestStore(conn Connection)
}

type Interactor struct {
	UserService domain.UserCase
	Output PresenterInput
}


func (interactor *Interactor) RequestIndex(conn Connection) {
	interactor.Output.PresentIndex(conn)
}

func (interactor *Interactor) RequestStore(conn Connection) {
	// Don't continue if required fields are missing.
	// validation
	for _, v := range []string{"firstname", "lastname", "email", "password"} {
		if len(conn.Request.FormValue(v)) == 0 {
			//call presenter present400
			interactor.Output.Present400(conn)
			return
		}
	}

	// Build the user from the form values.
	u := new(domain.User)
	u.FirstName = conn.Request.FormValue("firstname")
	u.LastName = conn.Request.FormValue("lastname")
	u.Email = conn.Request.FormValue("email")
	u.Password = conn.Request.FormValue("password")

	// Add the user to the database.
	err := interactor.UserService.CreateUser(u)
	if err != nil {
		//call presenter present500
		interactor.Output.Present500(conn, err)
	} else {
		//call presenter presentSuccess
		interactor.Output.PresentSuccessfulUserCreation(conn)
	}

}