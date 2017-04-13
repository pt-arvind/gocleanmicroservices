package login_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	//"tests(tobemoved)"
	"cmd/logic/handler/login"
	"domain"
	"lib/passhash"
	"lib/view"
	"logic"
	"repository"
)

// TestLoginIndex ensures the index function returns a 200 code.
func TestLoginIndex(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(login.Controller)
	h.ViewService = view.New("../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusOK)
}

// TestLoginStoreMissingRequiredField ensures required fields should be entered.
func TestLoginStoreMissingRequiredFields(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(login.Controller)
	h.UserService = logic.NewUserCase(
		repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	h.ViewService = view.New("../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
}

// TestLoginStoreAuthenticateOK ensures login can be successful.
func TestLoginStoreAuthenticateOK(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("email", "jdoe@example.com")
	r.Form.Add("password", "Pa$$w0rd")

	// Call the handler.
	h := new(login.Controller)
	h.UserService = logic.NewUserCase(
		repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	h.ViewService = view.New("../view", "tmpl")

	// Create a new user.
	u := new(domain.User)
	u.Email = "jdoe@example.com"
	u.Password = "Pa$$w0rd"
	h.UserService.CreateUser(u)

	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusOK)
}

// TestLoginStoreAuthenticateFail ensures login can fail.
func TestLoginStoreAuthenticateFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("email", "jdoe2@example.com")
	r.Form.Add("password", "BadPa$$w0rd")

	// Call the handler.
	h := new(login.Controller)
	h.UserService = logic.NewUserCase(
		repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	h.ViewService = view.New("../view", "tmpl")

	// Create a new user.
	u := new(domain.User)
	u.Email = "jdoe2@example.com"
	u.Password = "Pa$$w0rd"
	h.UserService.CreateUser(u)

	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusUnauthorized)
}
