package login

import (
	"net/http"
	"fmt"
	"github.com/pt-arvind/gocleanarchitecture/logic"
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

type PresenterInput interface {
	//Present404(conn Connection)
	//Present400(conn Connection)
	//Present401(conn Connection)
	//PresentSuccessfulLogin(conn Connection)
	//PresentIndex(conn Connection)
	logic.UserInteractorOutput
}

// should be functionally equivalent to http.ResponseWriter
type PresenterOutput interface {
	http.ResponseWriter
}

type Presenter struct {
	Output domain.ViewCase //TODO: should not be in domain
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

