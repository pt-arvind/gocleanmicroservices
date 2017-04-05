package login

import (
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

type InteractorInput interface {
	Request404(conn Connection)
	RequestIndex(conn Connection)
	RequestStore(conn Connection)
}

// should be equivalent to PresenterInput
//type InteractorOutput interface {
//	Present404()
//}

type Interactor struct {
	UserService domain.UserCase
	Output PresenterInput
}

func (interactor *Interactor) Request404(conn Connection) {
	interactor.Output.Present404(conn)
}

func (interactor *Interactor) RequestIndex(conn Connection) {
	interactor.Output.PresentIndex(conn)
}

func (interactor *Interactor) RequestStore(conn Connection) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(conn.Request.FormValue(v)) == 0 {
			interactor.Output.Present400(conn)
			return
		}
	}

	user := new(domain.User)

	user.Email = conn.Request.FormValue("email")
	user.Password = conn.Request.FormValue("password")

	err := interactor.UserService.Authenticate(user)

	if err != nil {
		interactor.Output.Present401(conn) //realistically would probably pass that error along
	} else {
		interactor.Output.PresentSuccessfulLogin(conn) // realistically, you'd want to have something here that would pass along the user you just made
	}
}

