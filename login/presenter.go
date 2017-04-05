package login

import (
	"net/http"
	"fmt"
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

type PresenterInput interface {
	Present404(conn Connection)
	Present400(conn Connection)
	Present401(conn Connection)
	PresentSuccessfulLogin(conn Connection)
	PresentIndex(conn Connection)
}

// should be functionally equivalent to http.ResponseWriter
type PresenterOutput interface {
	Header() http.Header
	Write([]byte) (int, error)
	WriteHeader(int)
}

type Presenter struct {
	//Output http.ResponseWriter
	ViewService domain.ViewCase
}

func (presenter *Presenter) Present400(conn Connection) {
	conn.Writer.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(conn.Writer, `<html>One or more required fields are missing. `+
		 	 `Click <a href="/">here</a> to try again.</html>`) // ideally you'd have a template file here and set a view model in the vars but for such a simple example, it's omitted
}

func (presenter *Presenter) Present401(conn Connection) {
	conn.Writer.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(conn.Writer, `<html>Login failed. `+
			 `Click <a href="/">here</a> to try again.</html>`)// ideally you'd have a template file here and set a view model in the vars but for such a simple example, it's omitted
}

func (presenter *Presenter) Present404(conn Connection) {
	conn.Writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(conn.Writer, "404 Page Not Found") // ideally you'd have a template file here and set a view model in the vars but for such a simple example, it's omitted
}

func (presenter *Presenter) PresentSuccessfulLogin(conn Connection) {
	fmt.Fprint(conn.Writer, "<html>Login successful!</html>")
}

func (presenter *Presenter) PresentIndex(conn Connection) {
	presenter.ViewService.SetTemplate("login/index")
	presenter.ViewService.Render(conn.Writer, conn.Request)
}