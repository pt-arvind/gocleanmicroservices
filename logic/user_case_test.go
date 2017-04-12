package logic_test

import (
	"errors"
	"testing"

	"github.com/pt-arvind/gocleanarchitecture/domain"
	"github.com/pt-arvind/gocleanarchitecture/lib/passhash"
	"github.com/pt-arvind/gocleanarchitecture/repository"
	"github.com/pt-arvind/gocleanarchitecture/logic"
)

//  BadHasher represents a password hashing system that always fails.
type BadHasher struct{}

// Hash returns a hashed string and an error.
func (s *BadHasher) Hash(password string) (string, error) {
	return "", errors.New("Forced error.")
}

// Match returns true if the hash matches the password.
func (s *BadHasher) Match(hash, password string) bool {
	return false
}

type MockUserInteractorOutput struct {
	DidCreate func(user domain.User) ()
	GotError func(err error) ()
	DidAuthenticate func(user domain.User) ()
	DidRetrieve func(user domain.User) ()
	IndexPassthru func() ()
}

func (o *MockUserInteractorOutput) Error(err error) {
	if o.GotError != nil {
		o.GotError(err)
	}
}

func (o *MockUserInteractorOutput) Authenticated(user domain.User) {
	if o.DidAuthenticate != nil {
		o.DidAuthenticate(user)
	}

}

func (o *MockUserInteractorOutput) UserCreated(user domain.User) {
	if o.DidCreate != nil {
		o.DidCreate(user)
	}
}

func (o *MockUserInteractorOutput) UserRetrieved(user domain.User) {
	if o.DidRetrieve != nil {
		o.DidRetrieve(user)
	}
}

func (o *MockUserInteractorOutput) Index() {
	if o.IndexPassthru != nil {
		o.IndexPassthru()
	}
}

type MockUser struct {
	FirstName string
	LastName string
	Email string
	Password string
}


var interactor logic.UserInteractorInput
var outputMock = &MockUserInteractorOutput{}

func setup() {
	factory := logic.NewInteractorFactory(repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))

	interactor = factory.NewUseCaseInteractor()
	interactor.SetOutput(outputMock)
}

func teardown() {
	outputMock.DidCreate = nil
	outputMock.GotError = nil
	outputMock.DidAuthenticate = nil
	outputMock.DidRetrieve = nil
	outputMock.IndexPassthru = nil
}


func TestSetOutput(t *testing.T) {
	factory := logic.NewInteractorFactory(repository.NewUserRepo(new(repository.MockService)),
		new(passhash.Item))

	interactor = factory.NewUseCaseInteractor()

	interactor.SetOutput(outputMock)


}

// TestCreateUser ensures user can be created.
func TestCreateUser(t *testing.T) {

	// Test user creation.
	firstname := "john"
	lastname := "doe"
	email := "john@example.com"
	password := "Pa$$w0rd"

	mockUser := MockUser{FirstName: firstname, LastName: lastname, Email: email, Password: password}

	setup()
	t.Run("successful user create", func(t *testing.T) {
		outputMock.DidCreate = func (user domain.User) {
			AssertEqual(t, user.FirstName, mockUser.FirstName)
			AssertEqual(t, user.LastName, mockUser.LastName)
			AssertEqual(t, user.Email, mockUser.Email)
			AssertEqual(t, user.Password, mockUser.Password)
		}
		outputMock.GotError = func (err error) {
			AssertFailure(t, "should not fail to create user...")
		}
		interactor.CreateUser(mockUser.FirstName,mockUser.LastName,mockUser.Email,mockUser.Password)
	})
	teardown()

	setup()
	t.Run("attempt to create duplicate user", func(t *testing.T) {
		// create user
		interactor.CreateUser(mockUser.FirstName,mockUser.LastName,mockUser.Email,mockUser.Password)

		outputMock.DidCreate = func (user domain.User) {
			AssertFailure(t, "should fail to create user...")
		}
		outputMock.GotError = func (err error) {
			AssertEqual(t, err, domain.ErrUserAlreadyExist)
		}
		// try to create user with same info
		interactor.CreateUser(mockUser.FirstName,mockUser.LastName,mockUser.Email,mockUser.Password)
	})
	teardown()

	//u := new(domain.User)
	//u.Email = "jdoe@example.com"
	//u.Password = "Pa$$w0rd"
	//err := s.CreateUser(u)

	//
	//AssertEqual(t, err, nil)
	//
	//// Test user creation fail.
	//err = s.CreateUser(u)
	//AssertEqual(t, err, domain.ErrUserAlreadyExist)
	//
	//// Test user query.
	//uTest, err := s.User("jdoe@example.com")
	//AssertEqual(t, err, nil)
	//AssertEqual(t, uTest.Email, "jdoe@example.com")
	//
	//// Test failed user query.
	//_, err = s.User("bademail@example.com")
	//AssertEqual(t, err, domain.ErrUserNotFound)

}
//
//// TestAuthenticate ensures user can authenticate.
func TestAuthenticate(t *testing.T) {

	firstname := "john"
	lastname := "doe"
	email := "john@example.com"
	password := "Pa$$w0rd"

	mockUser := MockUser{FirstName: firstname, LastName: lastname, Email: email, Password: password}

	setup()
	t.Run("successful login", func(t *testing.T) {
		outputMock.DidAuthenticate = func (user domain.User) {
			AssertEqual(t, user.FirstName, mockUser.FirstName)
			AssertEqual(t, user.LastName, mockUser.LastName)
			AssertEqual(t, user.Email, mockUser.Email)
			AssertEqual(t, user.Password, mockUser.Password)
		}
		outputMock.GotError = func (err error) {
			AssertFailure(t, "should not fail to login...")
		}
		interactor.CreateUser(mockUser.FirstName,mockUser.LastName,mockUser.Email,mockUser.Password)
		interactor.Authenticate(mockUser.Email, mockUser.Password)
	})
	teardown()

	setup()
	t.Run("imposter failed login", func(t *testing.T) {
		outputMock.DidAuthenticate = func (user domain.User) {
			AssertFailure(t, "should not authenticate...")
		}
		outputMock.GotError = func (err error) {
			AssertEqual(t, err, domain.ErrUserPasswordNotMatch)
		}
		interactor.CreateUser(mockUser.FirstName,mockUser.LastName,mockUser.Email,mockUser.Password)
		interactor.Authenticate(mockUser.Email, "wr0ngP@$$w0rD")
	})
	teardown()

	setup()
	t.Run("no such user", func(t *testing.T) {
		outputMock.DidAuthenticate = func (user domain.User) {
			AssertFailure(t, "should not authenticate...")
		}
		outputMock.GotError = func (err error) {
			AssertEqual(t, err, domain.ErrUserNotFound)
		}
		interactor.Authenticate(mockUser.Email, mockUser.Password)
	})
	teardown()
}
//
//// TestUserFailures ensures user fails properly.
//func TestUserFailures(t *testing.T) {
//	// Test user creation.
//	db := new(repository.MockService)
//	s := logic.NewUserCase(repository.NewUserRepo(db), new(passhash.Item))
//
//	db.WriteFail = true
//	db.ReadFail = true
//
//	u := new(domain.User)
//	u.Email = "ssmith@example.com"
//	u.Password = "Pa$$w0rd"
//	err := s.CreateUser(u)
//	AssertNotNil(t, err)
//
//	// Test user authentication.
//	err = s.Authenticate(u)
//	AssertNotNil(t, err)
//
//	// Test failed user query.
//	_, err = s.User("favalon@example.com")
//	AssertNotNil(t, err)
//
//	// Test failed user authentication.
//	u.Email = "bfranklin@example.com"
//	err = s.Authenticate(u)
//	AssertNotNil(t, err)
//}
//
//// TestBadHasherFailures ensures user fails properly.
//func TestBadHasherFailures(t *testing.T) {
//	// Test user creation.
//	db := new(repository.MockService)
//	s := logic.NewUserCase(repository.NewUserRepo(db), new(BadHasher))
//
//	u := new(domain.User)
//	u.Email = "ssmith@example.com"
//	u.Password = "Pa$$w0rd"
//	err := s.CreateUser(u)
//	AssertEqual(t, err, domain.ErrPasswordHash)
//}
