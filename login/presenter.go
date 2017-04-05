package login

import (
	"net/http"
	"fmt"
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

type PresenterInput interface {
	Present404(w http.ResponseWriter)
	Present400(w http.ResponseWriter)
	Present401(w http.ResponseWriter)
	PresentSuccessfulLogin(w http.ResponseWriter)
	PresentIndex(w http.ResponseWriter,r *http.Request)
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

func (presenter *Presenter) Present400(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, `<html>One or more required fields are missing. `+
		 	 `Click <a href="/">here</a> to try again.</html>`) // ideally you'd have a template file here and set a view model in the vars but for such a simple example, it's omitted
}

func (presenter *Presenter) Present401(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprint(w, `<html>Login failed. `+
			 `Click <a href="/">here</a> to try again.</html>`)// ideally you'd have a template file here and set a view model in the vars but for such a simple example, it's omitted
}

func (presenter *Presenter) Present404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page Not Found") // ideally you'd have a template file here and set a view model in the vars but for such a simple example, it's omitted
}

func (presenter *Presenter) PresentSuccessfulLogin(w http.ResponseWriter) {
	fmt.Fprint(w, "<html>Login successful!</html>")
}

func (presenter *Presenter) PresentIndex(w http.ResponseWriter, r *http.Request) {
	presenter.ViewService.SetTemplate("login/index")
	presenter.ViewService.Render(w, r)
}