package login

import (
	"net/http"
	"fmt"
	"cloudtamer/portal/cmd/webapp/logic"
	"cloudtamer/portal/domain"
	"cloudtamer/portal/cmd/webapp/adapter/viewport"

)

type PresenterInput interface {
	logic.UserInteractorOutput
}

// should be functionally equivalent to http.ResponseWriter
type PresenterOutput interface {
	http.ResponseWriter
}

type Presenter struct {
	Output viewport.ViewCase //TODO: should not be in domain
	Connection Connection // ugh :(
}

func (presenter *Presenter) Error(err error) {
	fmt.Fprint(presenter.Connection.Writer, "<html>Error!</html>")
}

func (presenter *Presenter) Authenticated(user domain.User) {
	fmt.Fprint(presenter.Connection.Writer, "<html>Login successful!</html>")
}

func (presenter *Presenter) UserCreated(user domain.User) {
	panic("shouldnt get called")
}

func (presenter *Presenter) UserRetrieved(user domain.User) {
	panic("shouldnt get called")
}

func (presenter *Presenter) Index() {
	presenter.Output.SetTemplate("login/index")
	presenter.Output.Render(presenter.Connection.Writer, presenter.Connection.Request)
}

