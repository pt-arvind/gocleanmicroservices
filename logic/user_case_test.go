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
	UserCreateSuccess func(user domain.User) ()
	GotError func(err error) ()
}

func (o *MockUserInteractorOutput) Error(err error) {
	o.GotError(err)
}

func (o *MockUserInteractorOutput) Authenticated(user domain.User) {
	panic("fuck 1")
}

func (o *MockUserInteractorOutput) UserCreated(user domain.User) {
	o.UserCreateSuccess(user)
}

func (o *MockUserInteractorOutput) UserRetrieved(user domain.User) {
	panic("fuck 3")
}

func (o *MockUserInteractorOutput) Index() {
	panic("fuck 4")
}

// TestCreateUser ensures user can be created.
func TestCreateUser(t *testing.T) {
	// Test user creation.
	factory := logic.NewInteractorFactory(repository.NewUserRepo(new(repository.MockService)),
					new(passhash.Item))

	interactor := factory.NewUseCaseInteractor()


	firstname := "john"
	lastname := "doe"
	email := "john@example.com"
	password := "Pa$$w0rd"

	mock := &MockUserInteractorOutput{}

	interactor.SetOutput(mock)


	// test create user successfully
	mock.UserCreateSuccess = func (user domain.User) {
		AssertEqual(t, user.FirstName, firstname)
		AssertEqual(t, user.LastName, lastname)
		AssertEqual(t, user.Email, email)
		AssertEqual(t, user.Password, password)
	}

	interactor.CreateUser(firstname,lastname,email,password)


	// test create user with same email (should fail)
	mock.GotError = func (err error) {
		AssertEqual(t, err, domain.ErrUserAlreadyExist)
	}
	interactor.CreateUser(firstname,lastname,email,password)

	// test bad user creation
	mock.GotError = func (err error) {
		AssertEqual(t, err, domain.ErrUserAlreadyExist)
	}

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



	// subject under test


	//

}
//
//// TestAuthenticate ensures user can authenticate.
//func TestAuthenticate(t *testing.T) {
//	// Test user creation.
//	s := logic.NewUserCase(repository.NewUserRepo(new(repository.MockService)),
//		new(passhash.Item))
//	u := new(domain.User)
//	u.Email = "ssmith@example.com"
//	u.Password = "Pa$$w0rd"
//	err := s.CreateUser(u)
//	AssertEqual(t, err, nil)
//
//	// Test user authentication.
//	err = s.Authenticate(u)
//	AssertEqual(t, err, nil)
//
//	// Test failed user authentication.
//	u.Password = "BadPa$$w0rd"
//	err = s.Authenticate(u)
//	AssertEqual(t, err, domain.ErrUserPasswordNotMatch)
//
//	// Test failed user authentication.
//	u.Email = "bfranklin@example.com"
//	err = s.Authenticate(u)
//	AssertEqual(t, err, domain.ErrUserNotFound)
//}
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
