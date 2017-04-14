package logic

import (
	"domain"
	"fmt"
)


type UserInteractorFactory struct {

}

func (f UserInteractorFactory) NewUseCaseInteractor(userRepo UserRepoInput) UserInteractorInput {
	interactor := new(UserInteractor)
	interactor.repo = userRepo
	userRepo.SetOutput(interactor)

	return interactor
}


// NewUserCase returns the service for managing users.
func NewInteractorFactory() *UserInteractorFactory {
	s := new(UserInteractorFactory)
	return s
}

type UserRepoInput interface {
	FindByID(int)
	Store(user domain.User)
	AllUsers()
	SetOutput(UserRepoOutput)
}

type UserRepoOutput interface {
	Found(user domain.User)
	Stored(user domain.User)
	Users(users []domain.User)
	Error(error)
}

type UserInteractorInput interface {
	SetOutput(output UserInteractorOutput)
	Error(err error)

	GetAllUsers()

	// outputs a user by email
	GetUser(id int)

	// CreateUser creates a new user and outputs it
	CreateUser(domain.User)
}

type UserInteractorOutput interface {
	Error(err error)
	UserStored(user domain.User)
	UserFound(user domain.User)
	AllUsers(users []domain.User)
}

type UserInteractor struct {
	output UserInteractorOutput
	repo UserRepoInput
}

/* UserRepoOutput conformance  */
func (i UserInteractor) Found(user domain.User) {
	i.output.UserFound(user)
}

func (i UserInteractor) Stored(user domain.User) {
	i.output.UserStored(user)
}

func (i UserInteractor) Error(err error) {
	i.output.Error(err)
}

func (i UserInteractor) Users(users []domain.User) {
	fmt.Println("GOT ALL USERS TIME TO PRESENT")
	i.output.AllUsers(users)
}


/* UserInteractorInput conformance */
func (i *UserInteractor) SetOutput(output UserInteractorOutput) {
	i.output = output
}

func (i *UserInteractor) GetAllUsers() {
	fmt.Println("GETALLUSERS")
	i.repo.AllUsers()
}

func (i UserInteractor) GetUser(id int) {
	i.repo.FindByID(id)
}

// CreateUser creates a new user and outputs it
func (i UserInteractor) CreateUser(user domain.User) {
	i.repo.Store(user)
}
