package boot

import (
	//"cmd/webapp/adapter/repository"
	//"cmd/webapp/adapter/viewport"
	//"cmd/webapp/infrastructure/jsondb"
	//"cmd/webapp/infrastructure/passhash"
	//"cmd/webapp/infrastructure/view"
	//"cmd/webapp/logic"

	"cmd/dbservice/infrastructure/viewjson"
	"cmd/dbservice/infrastructure/jsondb"
	"cmd/dbservice/logic"
)

// Service represents all the services that the application uses.
type Service struct {
	Database    *jsondb.Client
	UserService logic.UserInteractorFactory
	ViewService view.JSONViewFactory
}

// RegisterServices sets up each service and returns the container for all
// the services.
func RegisterServices() *Service {
	s := new(Service)

	// Initialize the clients.
	s.Database = jsondb.NewClient("db.json")

	s.UserService = *logic.NewInteractorFactory()
	s.ViewService = view.JSONViewFactory{}

	return s
}
