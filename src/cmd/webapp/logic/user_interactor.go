package logic

import (
	"domain"
)

// UserInteractor represents a service for managing users.
type UserInteractorInput interface {
	SetOutput(output UserInteractorOutput)
	Error(err error)
	Authenticate(email string, password string)
	User(email string)
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


type UserRepoInput interface {
	FindByID(id int, output func(user *domain.User, err error))
	Store(item domain.User, output func(user *domain.User, err error))
	AllUsers(output func(users []domain.User, err error))
}

type UserInteractor struct {
	userRepo domain.UserRepo
	passhash domain.PasshashCase
	output 	 UserInteractorOutput
}

type UserInteractorFactory struct {
	passhash domain.PasshashCase
}

func (factory UserInteractorFactory) NewUseCaseInteractor(repo domain.UserRepo) UserInteractorInput {
	interactor := new(UserInteractor)
	interactor.userRepo = repo
	interactor.passhash = factory.passhash

	return interactor
}


// NewUserCase returns the service for managing users.
func NewInteractorFactory(passhash domain.PasshashCase) *UserInteractorFactory {
	s := new(UserInteractorFactory)
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
	interactor.userRepo.FindByEmail(email, func(user *domain.User, err error) {
		if err != nil {
			interactor.output.Error(domain.ErrUserNotFound)
			return
		}

		// If passwords match.
		if interactor.passhash.Match(user.Password, password) {
			user.Password = password //unhashed pass to keep consistency
			interactor.output.Authenticated(*user)
			return
		}

		interactor.output.Error(domain.ErrUserPasswordNotMatch)
	})

}

// User returns a user by email.
func (interactor *UserInteractor) User(email string) {
	interactor.userRepo.FindByEmail(email, func(user *domain.User, err error) {
		if err != nil {
			interactor.output.Error(domain.ErrUserNotFound)
			return
		}

		interactor.output.UserRetrieved(*user)
	})


}

// CreateUser creates a new user.
func (interactor *UserInteractor) CreateUser(firstName string, lastName string, email string, password string) {

	interactor.userRepo.FindByEmail(email, func(user *domain.User, err error) {
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


		interactor.userRepo.Store(*item, func(user *domain.User, err error) {
			// Restore the password.
			item.Password = passOld

			if err != nil {
				interactor.output.Error(err)
				return
			}

			interactor.output.UserCreated(*item)

		})

	})




}
