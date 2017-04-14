package user

import (
	"net/http"
	"cmd/dbservice/logic"
	"domain"
	"cmd/dbservice/adapter/viewport"
)

type PresenterInput interface {
	logic.UserInteractorOutput
}

type PresenterOutput interface {
	http.ResponseWriter
}

type Presenter struct {
	Connection Connection
	Output viewport.Viewport
}

func (p *Presenter) Error(err error) {
	p.Output.Render(p.Connection.Writer, nil, err)
}

func (p *Presenter) UserStored(user domain.User) {
	users := []domain.User{user}
	p.Output.Render(p.Connection.Writer, users, nil)
}

func (p *Presenter) UserFound(user domain.User) {
	users := []domain.User{user}
	p.Output.Render(p.Connection.Writer, users, nil)
}

func (p *Presenter) AllUsers(users []domain.User) {
	p.Output.Render(p.Connection.Writer, users, nil)
}