package boot

import (
	"net/http"

	"github.com/pt-arvind/gocleanarchitecture/cmd/webapp/adapter/handler/login"
	"github.com/pt-arvind/gocleanarchitecture/cmd/webapp/adapter/handler/register"
	//"fmt"
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
	//controller.UserInteractor = s.UserInteractor
	//controller.ViewService = s.ViewService


	//interactor := new(login.Interactor)
	//interactor.UserInteractor = s.UserService

	//:_(
	userService := s.UserService

	//val := reflect.ValueOf(userService)
	//if val.Kind() == reflect.Ptr {
	//	val = reflect.Indirect(val)
	//}
	//interactor := reflect.New(val.Type()).Interface().(logic.UserInteractorInput)

	interactor := userService.NewUseCaseInteractor()

	presenter := new(login.Presenter)
	presenter.Output = s.ViewService

	// hook up the flow, interactor -> presenter
	interactor.SetOutput(presenter)

	// controller -> interactor
	controller.Output = interactor
	controller.Presenter = presenter // :( this is so that we can set the connection on the presenter as it passes through


	// Load routes.
	mux.HandleFunc("/", controller.Route)
}

// AddRegister registers the register handlers.
func (s *Service) AddRegister(mux *http.ServeMux) {
	// Create handler.
	//controller := new(register.Controller)
	//
	//interactor := new(register.Interactor)
	//interactor.UserService = s.UserService
	//
	//
	//presenter := new(register.Presenter)
	//presenter.ViewService = s.ViewService
	//
	//// hook up the flow, interactor -> presenter
	//interactor.Output = presenter
	//
	//// controller -> interactor
	//controller.Output = interactor


	// CAUTION: this stuff has to be set up in this way because of pass-by-value vs pass-by-reference semantics!
	userService := s.UserService

	//val := reflect.ValueOf(userService)
	//fmt.Println(val)
	//fmt.Println(reflect.ValueOf(val))
	//
	//
	//if val.Kind() == reflect.Ptr {
	//	val = reflect.Indirect(val)
	//}
	//fmt.Println("=============================================")
	//fmt.Println(userService)
	//fmt.Println(val)
	//fmt.Println(val.Type())
	//fmt.Println(reflect.ValueOf(val))
	//fmt.Println("interface time")
	//fmt.Println(reflect.ValueOf(val).Interface())
	//
	//cval := reflect.ValueOf(val)
	//nval := cval.Interface()
	//
	//fmt.Println(nval)
	//
	//
	////var _ logic.UserInteractorInput = (*logic.UserInteractor)(nil) // compiles because logic.UserInteractor conforms to UserInteractorInput....
	//
	//aval := &logic.UserInteractor{}
	//var bval logic.UserInteractorInput = aval
	//
	//fmt.Println(bval)


	interactor := userService.NewUseCaseInteractor() //(reflect.ValueOf(val).Interface()).(logic.UserInteractorInput)//reflect.New(val.Type()).Interface().(logic.UserInteractorInput)
	//fmt.Println(interactor)

	controller := new(register.Controller)

	presenter := new(register.Presenter)
	presenter.Output = s.ViewService

	interactor.SetOutput(presenter)

	controller.Output = interactor
	controller.Presenter = presenter


	// Assign services.
	//h.UserInteractor = s.UserInteractor
	//h.ViewService = s.ViewService

	// Load routes.
	mux.HandleFunc("/register", controller.Route)
}
