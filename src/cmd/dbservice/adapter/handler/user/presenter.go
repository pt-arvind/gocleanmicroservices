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

type Presenter struct {
	Connection http.ResponseWriter
	Output viewport.Viewport
}

func (p *Presenter) Error(err error) {
	p.Output.Render(p.Connection, nil, err)
}

func (p *Presenter) UserStored(user domain.User) {
	users := []domain.User{user}
	p.Output.Render(p.Connection, users, nil)
}

func (p *Presenter) UserFound(user domain.User) {
	users := []domain.User{user}
	p.Output.Render(p.Connection, users, nil)
}

func (p *Presenter) AllUsers(users []domain.User) {
	p.Output.Render(p.Connection, users, nil)
}