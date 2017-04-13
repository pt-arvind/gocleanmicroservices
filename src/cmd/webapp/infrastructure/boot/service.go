package boot

import (
	"cloudtamer/portal/cmd/webapp/infrastructure/passhash"
	"cloudtamer/portal/cmd/webapp/infrastructure/view"
	"cloudtamer/portal/cmd/webapp/infrastructure/jsondb"
	"cloudtamer/portal/cmd/webapp/adapter/repository"
	"cloudtamer/portal/cmd/webapp/logic"
	"cloudtamer/portal/cmd/webapp/adapter/viewport"
)

// Service represents all the services that the application uses.
type Service struct {
	UserService logic.UserInteractorFactory
	ViewService viewport.ViewCase
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices() *Service {
	s := new(Service)

	// Initialize the clients.
	db := jsondb.NewClient("db.json") //FIXME: this path is not putting it in /webapp/infrastructure/jsondb but rather in /webapp

	// Store all the services for the application.
	s.UserService = *logic.NewInteractorFactory(
		repository.NewUserRepo(db),
		new(passhash.Item))
	s.ViewService = view.New("/infrastructure/html/", "tmpl") //FIXME: code smell, the path here is relative to the view package not THIS package!

	return s
}