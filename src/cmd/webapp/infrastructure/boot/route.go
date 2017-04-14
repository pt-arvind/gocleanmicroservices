package boot

import (
	"net/http"

	"cmd/webapp/adapter/handler/login"
	"cmd/webapp/adapter/handler/register"
	"cmd/webapp/adapter/repository"
)

// LoadRoutes returns a handler with all the routes.
func (s *Service) LoadRoutes() http.Handler {
	// Create the mux.
	h := http.NewServeMux()

	// Register the pages.
	s.AddLogin(h)
	s.AddRegister(h)

	// Return the handler.
	return h
}

// AddLogin registers the login handlers.
func (s *Service) AddLogin(mux *http.ServeMux) {
	// Create handler.
	controller := new(login.Controller)

	userService := s.UserService

	repo := repository.NewUserRepo(s.DBService)
	interactor := userService.NewUseCaseInteractor(repo)

	presenter := new(login.Presenter)
	presenter.Output = s.ViewService

	// hook up the flow, interactor -> presenter
	interactor.SetOutput(presenter)

	// tests(tobemoved) -> interactor
	controller.Output = interactor
	controller.Presenter = presenter // :( this is so that we can set the connection on the presenter as it passes through


	// Load routes.
	mux.HandleFunc("/", controller.Route)
}

// AddRegister registers the register handlers.
func (s *Service) AddRegister(mux *http.ServeMux) {
	// CAUTION: this stuff has to be set up in this way because of pass-by-value vs pass-by-reference semantics!
	userService := s.UserService

	repo := repository.NewUserRepo(s.DBService)
	interactor := userService.NewUseCaseInteractor(repo)

	controller := new(register.Controller)

	presenter := new(register.Presenter)
	presenter.Output = s.ViewService

	interactor.SetOutput(presenter)

	controller.Output = interactor
	controller.Presenter = presenter

	// Load routes.
	mux.HandleFunc("/register", controller.Route)
}
