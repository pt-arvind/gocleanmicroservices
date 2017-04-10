package boot

import (
	"net/http"

	"github.com/pt-arvind/gocleanarchitecture/login"
	"github.com/pt-arvind/gocleanarchitecture/register"
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

	// Assign services.
	//controller.UserService = s.UserService
	//controller.ViewService = s.ViewService


	interactor := new(login.Interactor)
	interactor.UserService = s.UserService


	presenter := new(login.Presenter)
	presenter.ViewService = s.ViewService

	// hook up the flow, interactor -> presenter
	interactor.Output = presenter

	// controller -> interactor
	controller.Output = interactor


	// Load routes.
	mux.HandleFunc("/", controller.Index)
}

// AddRegister registers the register handlers.
func (s *Service) AddRegister(mux *http.ServeMux) {
	// Create handler.
	controller := new(register.Controller)

	interactor := new(register.Interactor)
	interactor.UserService = s.UserService


	presenter := new(register.Presenter)
	presenter.ViewService = s.ViewService

	// hook up the flow, interactor -> presenter
	interactor.Output = presenter

	// controller -> interactor
	controller.Output = interactor

	// Assign services.
	//h.UserService = s.UserService
	//h.ViewService = s.ViewService

	// Load routes.
	mux.HandleFunc("/register", controller.Index)
}
