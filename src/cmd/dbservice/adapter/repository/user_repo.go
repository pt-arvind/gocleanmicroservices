package repository

//adapter layer

import (
	"domain"
	"errors"
	"cmd/dbservice/logic"
)


var (
	// ErrUserNotFound is when the user does not exist.
	ErrUserNotFound = errors.New("User not found.")
	// ErrUserAlreadyExist is when the user already exists.
	ErrUserAlreadyExists = errors.New("User already exists.")
)

// Service represents a service for interacting with the database.
type ServiceInput interface {
	//Read() error
	//Write() error
	//Records() []domain.User
	//AddRecord(domain.User)

	Load(output func(err error))
	Save(output func(err error))
	Records(output func([]domain.User))
	AddRecord(user domain.User, output func(user domain.User))
}

// UserRepo represents a service for storage of users.
type UserRepo struct {
	client ServiceInput
	output logic.UserRepoOutput
}

// NewUserRepo returns the service for storage of users.
func NewUserRepo(client ServiceInput) *UserRepo {
	s := new(UserRepo)
	s.client = client
	//s.client.SetOutput(s)
	return s
}

func (s *UserRepo) FindByID(id int){
	s.client.Records(func (users []domain.User){
		//check
		for _, user := range users {
			if user.ID == id {
				s.output.Found(user)
				return
			}

		}
		s.output.Error(domain.ErrUserNotFound)
	})
}

// Store adds a user.
func (s *UserRepo) Store(item domain.User){
	s.client.AddRecord(item, func(user domain.User) {
		s.client.Save(func(err error) {
			if err != nil {
				s.output.Error(err)
			} else {
				s.output.Stored(user)
			}
		})
	})
}

func (s *UserRepo) SetOutput(out logic.UserRepoOutput) {
	s.output = out
}

func (s *UserRepo) AllUsers() {
	s.client.Records(func (users []domain.User) {
		s.output.Users(users)
	})
}