package register

import (
	"net/http"
	"domain"
	"cmd/webapp/adapter/viewport"
	"fmt"
)

type PresenterInput interface {
	Present400(conn Connection)
	Present500(conn Connection, err error)
	PresentSuccessfulUserCreation(conn Connection)
	PresentIndex(conn Connection)
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
	//fmt.Println("ERROR")
	fmt.Fprint(presenter.Connection.Writer, "<html>Error!</html>")
}

func (presenter *Presenter) Authenticated(user domain.User) {
	panic("shouldnt get called")
}

func (presenter *Presenter) UserCreated(user domain.User) {
	//fmt.Println("CREATED!")
	fmt.Fprint(presenter.Connection.Writer, "<html>User Creation successful!</html>")
}

func (presenter *Presenter) UserRetrieved(user domain.User) {
	//fmt.Println("RETRIEVED!")
	fmt.Fprint(presenter.Connection.Writer, "<html>User retrieve successful!</html>")
}

func (presenter *Presenter) Index() {
	//fmt.Println("INDEX")
	presenter.Output.SetTemplate("register/index")
	presenter.Output.Render(presenter.Connection.Writer, presenter.Connection.Request)
}