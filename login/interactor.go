package login

import (
	"net/http"
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

type InteractorInput interface {
	Request404(w http.ResponseWriter, r *http.Request)
	RequestIndex(w http.ResponseWriter, r *http.Request)
	RequestStore(w http.ResponseWriter, r *http.Request)
}

// should be equivalent to PresenterInput
//type InteractorOutput interface {
//	Present404()
//}

type Interactor struct {
	UserService domain.UserCase
	Output PresenterInput
}

func (interactor *Interactor) Request404(w http.ResponseWriter, r *http.Request) {
	interactor.Output.Present404(w)
}

func (interactor *Interactor) RequestIndex(w http.ResponseWriter, r *http.Request) {
	interactor.Output.PresentIndex(w,r)
}

func (interactor *Interactor) RequestStore(w http.ResponseWriter, r *http.Request) {
	// Don't continue if required fields are missing.
	for _, v := range []string{"email", "password"} {
		if len(r.FormValue(v)) == 0 {
			interactor.Output.Present400(w)
			return
		}
	}

	user := new(domain.User)

	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")

	err := interactor.UserService.Authenticate(user)

	if err != nil {
		interactor.Output.Present401(w) //realistically would probably pass that error along
	} else {
		interactor.Output.PresentSuccessfulLogin(w) // realistically, you'd want to have something here that would pass along the user you just made
	}
}

