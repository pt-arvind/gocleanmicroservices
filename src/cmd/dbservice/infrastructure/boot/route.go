package boot

import (
	"net/http"

	"cmd/dbservice/adapter/handler/user"
	"cmd/dbservice/adapter/repository"
	"github.com/blue-jay/core/router"
)

// LoadRoutes returns a handler with all the routes.
func (s *Service) LoadRoutes() http.Handler {
	// Create the mux.
	//h := http.NewServeMux()

	// Register the pages.
	//s.AddLogin(h)
	//s.AddRegister(h)

	handler := router.Instance()

	s.AddUserEndpoint()


	// Return the handler.
	return handler
}

type ParamExtractor struct {
}
func (p ParamExtractor) Param(r *http.Request, name string) string {
	return router.Param(r, name)
}

func (s *Service) AddUserEndpoint() {
	controller := new(user.Controller)

	userService := s.UserService
	viewService := s.ViewService
	repo := repository.NewUserRepo(s.Database)

	interactor := userService.NewUseCaseInteractor(repo)
	viewport := viewService.NewJSONViewport()

	presenter := new(user.Presenter)
	presenter.Output = viewport

	// hook up the flow, interactor -> presenter
	interactor.SetOutput(presenter)

	controller.Output = interactor
	controller.ParamExtractor = ParamExtractor{}


	// Load routes.
	router.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		presenter.Connection = w
		controller.Create(r)

	})
	router.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		presenter.Connection = w
		controller.Index(r)
	})
	router.Get("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		presenter.Connection = w
		controller.Show(r)
	})

	//controller.Route()
}
//
//// AddLogin registers the login handlers.
//func (s *Service) AddLogin(mux *http.ServeMux) {
//	// Create handler.
//	controller := new(login.Controller)
//
//	userService := s.UserService
//
//	interactor := userService.NewUseCaseInteractor()
//
//	presenter := new(login.Presenter)
//	presenter.Output = s.ViewService
//
//	// hook up the flow, interactor -> presenter
//	interactor.SetOutput(presenter)
//
//	// tests(tobemoved) -> interactor
//	controller.Output = interactor
//	controller.Presenter = presenter // :( this is so that we can set the connection on the presenter as it passes through
//
//
//	// Load routes.
//	mux.HandleFunc("/", controller.Route)
//}
//
//// AddRegister registers the register handlers.
//func (s *Service) AddRegister(mux *http.ServeMux) {
//	// CAUTION: this stuff has to be set up in this way because of pass-by-value vs pass-by-reference semantics!
//	userService := s.UserService
//
//	interactor := userService.NewUseCaseInteractor()
//
//	controller := new(register.Controller)
//
//	presenter := new(register.Presenter)
//	presenter.Output = s.ViewService
//
//	interactor.SetOutput(presenter)
//
//	controller.Output = interactor
//	controller.Presenter = presenter
//
//	// Load routes.
//	mux.HandleFunc("/register", controller.Route)
//}
