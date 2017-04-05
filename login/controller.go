package login

import (
	"net/http"
	//"crypto/tls"
	//"net/url"
	//"io"
	//"mime/multipart"
	//"context"
)

// no need for these, because the controller is going to configure both presenter and interactor
//type ControllerInput interface {
//
//}
//
//type ControllerOutput interface {
//
//}

// LoginHandler represents the services required for this controller.
type Controller struct {
	//UserService domain.UserCase
	//ViewService domain.ViewCase
	Interactor  	InteractorInput
}

//type RequestInterface interface {
//	Method() string
//	URL() *url.URL
//
//}

//type Request struct {
//	Method string
//
//	URL *url.URL
//
//	Proto      string // "HTTP/1.0"
//	ProtoMajor int    // 1
//	ProtoMinor int    // 0
//	Header http.Header
//
//
//	Body io.ReadCloser
//
//	GetBody func() (io.ReadCloser, error)
//
//
//	ContentLength int64
//
//	TransferEncoding []string
//
//	Close bool
//	Host string
//
//	Form url.Values
//
//	PostForm url.Values
//
//	MultipartForm *multipart.Form
//
//	Trailer http.Header
//	RemoteAddr string
//	RequestURI string
//
//	TLS *tls.ConnectionState
//	Cancel <-chan struct{}
//
//	Response *http.Response
//	ctx context.Context
//}

// potentially use a struct to encapsulate http.ResponseWriter and http.Request in interfaces that you would then pass around
type Request struct {

	Writer PresenterOutput
}




type ViewModel struct {

}


// Index displays the logon screen.
func (h *Controller) Index(w http.ResponseWriter, r *http.Request) {
	// Handle 404.
	if r.URL.Path != "/" { // FIXME: will typically be handled by a router, so it's OK for this logic to be in here for now
		h.Interactor.Request404(w,r)
	} else if r.Method == "POST" { 	// FIXME: will typically be handled by a router, so it's OK for this logic to be in here for now
		// call store on interactor
		h.Interactor.RequestStore(w,r)
	} else {
		h.Interactor.RequestIndex(w,r)
	}
}

// Store handles the submission of the login information.
//func (h *Controller) Store(w http.ResponseWriter, r *http.Request) {
//}
