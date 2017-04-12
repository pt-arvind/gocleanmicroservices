package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/pt-arvind/gocleanarchitecture/tests(tobemoved)"
	"github.com/pt-arvind/gocleanarchitecture/lib/passhash"
	"github.com/pt-arvind/gocleanarchitecture/lib/view"
	"github.com/pt-arvind/gocleanarchitecture/repository"
	"github.com/pt-arvind/gocleanarchitecture/logic"
)

// TestRegisterIndex ensures the index function returns a 200 code.
func TestRegisterIndex(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(tests.RegisterHandler)
	h.ViewService = view.New("../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusOK)
}

// TestRegisterStoreCreateOK ensures register can be successful.
func TestRegisterStoreCreateOK(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("firstname", "John")
	r.Form.Add("lastname", "Doe")
	r.Form.Add("email", "jdoe@example.com")
	r.Form.Add("password", "Pa$$w0rd")

	// Call the handler.
	h := new(tests.RegisterHandler)
	h.UserService = logic.NewUserCase(
		repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	h.ViewService = view.New("../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusCreated)

	// Fail on duplicate creation.
	w = httptest.NewRecorder()
	h.Index(w, r)
	AssertEqual(t, w.Code, http.StatusInternalServerError)
}

// TestRegisterStoreCreateNoFieldFail ensures register can fail with no fields.
func TestRegisterStoreCreateNoFieldFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the handler.
	h := new(tests.RegisterHandler)
	h.UserService = logic.NewUserCase(
		repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	h.ViewService = view.New("../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
}

// TestRegisterStoreCreateOneMissingFieldFail ensures register can fail with one missing
// field.
func TestRegisterStoreCreateOneMissingFieldFail(t *testing.T) {
	// Set up the request.
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the request body.
	val := url.Values{}
	r.Form = val
	r.Form.Add("firstname", "John")
	//r.Form.Add("lastname", "Doe")
	r.Form.Add("email", "jdoe@example.com")
	r.Form.Add("password", "Pa$$w0rd")

	// Call the handler.
	h := new(tests.RegisterHandler)
	h.UserService = logic.NewUserCase(
		repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))
	h.ViewService = view.New("../view", "tmpl")
	h.Index(w, r)

	// Check the output.
	AssertEqual(t, w.Code, http.StatusBadRequest)
}
