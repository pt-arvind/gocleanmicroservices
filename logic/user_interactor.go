package logic

import (
	"github.com/pt-arvind/gocleanarchitecture/domain"
)

// UserInteractor represents a service for managing users.
type UserInteractorInput interface {
	SetOutput(output UserInteractorOutput)
	Error(err error)
	// Authenticate outputs an error if the email and password don't match.
	Authenticate(email string, password string)
	// outputs a user by email
	User(email string)
	// CreateUser creates a new user and outputs it
	CreateUser(firstName string, lastName string, email string, password string)
	Index()
}

type UserInteractorOutput interface {
	Error(err error)
	Authenticated(user domain.User)
	UserCreated(user domain.User)
	UserRetrieved(user domain.User)
	Index()
}

type UserInteractor struct {
	userRepo domain.UserRepo
	passhash domain.PasshashCase
	output 	 UserInteractorOutput
}

// NewUserCase returns the service for managing users.
func NewInteractor(repo domain.UserRepo, passhash domain.PasshashCase) *UserInteractor {
	s := new(UserInteractor)
	s.userRepo = repo
	s.passhash = passhash
	return s
}

// setters
func (interactor *UserInteractor) SetOutput(output UserInteractorOutput) {
	interactor.output = output
}

// passthrough
func (interactor *UserInteractor) Error(err error) {
	interactor.output.Error(err)
}

func (interactor *UserInteractor) Index() {
	interactor.output.Index()
}

// Authenticate outputs an error if the email and password don't match.
func (interactor *UserInteractor) Authenticate(email string, password string) {
	item, err := interactor.userRepo.FindByEmail(email)
	if err != nil {
		//return domain.ErrUserNotFound
		interactor.output.Error(domain.ErrUserNotFound)
		return
	}

	// If passwords match.
	if interactor.passhash.Match(item.Password, password) {
		item.Password = password //unhashed pass to keep consistency
		interactor.output.Authenticated(*item)
		//return nil
		return
	}

	//return domain.ErrUserPasswordNotMatch
	interactor.output.Error(domain.ErrUserPasswordNotMatch)
}

// User returns a user by email.
func (interactor *UserInteractor) User(email string) {
	item, err := interactor.userRepo.FindByEmail(email)
	if err != nil {
		interactor.output.Error(domain.ErrUserNotFound)
		return
		//return item, domain.ErrUserNotFound
	}

	//return item, nil
	interactor.output.UserRetrieved(*item)
}

// CreateUser creates a new user.
func (interactor *UserInteractor) CreateUser(firstName string, lastName string, email string, password string) {
	_, err := interactor.userRepo.FindByEmail(email)
	if err == nil {
		//return domain.ErrUserAlreadyExist
		interactor.output.Error(domain.ErrUserAlreadyExist)
		return
	}


	passNew, err := interactor.passhash.Hash(password)
	if err != nil {
		//return domain.ErrPasswordHash
		interactor.output.Error(domain.ErrPasswordHash)
		return
	}

	// Swap the password.
	passOld := password
	item := &domain.User{FirstName: firstName, LastName: lastName, Email: email, Password: passNew}

	err = interactor.userRepo.Store(item)

	// Restore the password.
	item.Password = passOld

	if err != nil {
		interactor.output.Error(err)
		return
	}
	interactor.output.UserCreated(*item)
	//return err
}
