package viewport

import (
	"net/http"
	"domain"
)

type Viewport interface {
	Render(w http.ResponseWriter, users []domain.User, err error)
}
