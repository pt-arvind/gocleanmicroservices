package boot

import (
	"github.com/pt-arvind/gocleanarchitecture/domain"
	"github.com/pt-arvind/gocleanarchitecture/lib/passhash"
	"github.com/pt-arvind/gocleanarchitecture/cmd/webapp/adapter/view"
	"github.com/pt-arvind/gocleanarchitecture/cmd/webapp/infrastructure/jsondb"
	"github.com/pt-arvind/gocleanarchitecture/cmd/webapp/adapter/repository"
	"github.com/pt-arvind/gocleanarchitecture/cmd/webapp/logic"
)

// Service represents all the services that the application uses.
type Service struct {
	UserService logic.UserInteractorFactory
	ViewService domain.ViewCase
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