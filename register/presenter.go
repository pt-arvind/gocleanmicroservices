package register

import (
	"net/http"
	"github.com/pt-arvind/gocleanarchitecture/domain"
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
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

type Presenter struct {
	ViewService domain.ViewCase
}

func (presenter *Presenter) PresentIndex(conn Connection) {
	presenter.ViewService.SetTemplate("register/index")
	presenter.ViewService.Render(conn.Writer, conn.Request)
}

func (presenter *Presenter) Present400(conn Connection) {
	conn.Writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(conn.Writer, `<html>One or more required fields are missing. `+
				`Click <a href="/register">here</a> to try again.</html>`)
}

func (presenter *Presenter) Present500(conn Connection, err error) {
	conn.Writer.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(conn.Writer, err)
}

func (presenter *Presenter) PresentSuccessfulUserCreation(conn Connection) {
	conn.Writer.WriteHeader(http.StatusCreated)
	fmt.Fprint(conn.Writer, `<html>User created. `+
				`Click <a href="/">here</a> to login.</html>`)
}