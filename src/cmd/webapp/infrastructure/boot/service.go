package boot

import (
	"cmd/webapp/adapter/viewport"
	"cmd/webapp/infrastructure/passhash"
	"cmd/webapp/infrastructure/view"
	"cmd/webapp/logic"
	"cmd/webapp/infrastructure/dbservice"
	"net/http"
	"time"
)

// Service represents all the services that the application uses.
type Service struct {
	UserService logic.UserInteractorFactory
	ViewService viewport.ViewCase
	DBService dbservice.Client
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices() *Service {
	s := new(Service)

	// Initialize the clients.

	s.DBService = dbservice.Client{Client:http.Client{Timeout: 10 * time.Second}}

	// Store all the services for the application.
	s.UserService = *logic.NewInteractorFactory(new(passhash.Item))
	s.ViewService = view.New("infrastructure/html/", "tmpl")

	return s
}
