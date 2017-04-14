package boot

import (
	"net/http"

	"cmd/dbservice/adapter/handler/user"
	"cmd/dbservice/adapter/repository"
	"github.com/blue-jay/core/router"
)

// LoadRoutes returns a handler with all the routes.
func (s *Service) LoadRoutes() http.Handler {

	handler := router.Instance()

	s.AddUserEndpoint()

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
}