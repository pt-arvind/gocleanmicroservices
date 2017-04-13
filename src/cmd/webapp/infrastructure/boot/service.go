package boot

import (
	"cmd/webapp/adapter/repository"
	"cmd/webapp/adapter/viewport"
	"cmd/webapp/infrastructure/jsondb"
	"cmd/webapp/infrastructure/passhash"
	"cmd/webapp/infrastructure/view"
	"cmd/webapp/logic"
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
	db := jsondb.NewClient("db.json")

	// Store all the services for the application.
	s.UserService = *logic.NewInteractorFactory(
		repository.NewUserRepo(db),
		new(passhash.Item))
	s.ViewService = view.New("infrastructure/html/", "tmpl")

	return s
}
