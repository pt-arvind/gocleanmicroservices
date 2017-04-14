package repository

//adapter layer

import (
	"domain"
)

// Service represents a service for interacting with the database service
type ServiceInput interface {
	Records(output func(users []domain.User, err error))
	AddRecord(user domain.User, output func(user *domain.User, err error))
}

// UserRepo represents a service for storage of users.
type UserRepo struct {
	client ServiceInput
}

// NewUserRepo returns the service for storage of users.
func NewUserRepo(client ServiceInput) *UserRepo {
	s := new(UserRepo)
	s.client = client
	return s
}

func (s *UserRepo) FindByEmail(email string, output func(user *domain.User, err error)){
	s.client.Records(func (users []domain.User, err error){
		if err != nil {
			output(nil, err)
			return
		}
		//check
		for _, user := range users {
			if user.Email == email {
				output(&user, nil)
				return
			}

		}
		output(nil, domain.ErrUserNotFound)
	})
}

// Store adds a user.
func (s *UserRepo) Store(item domain.User, output func(user *domain.User, err error)){
	s.client.AddRecord(item, func(user *domain.User, err error) {
		if err != nil {
			output(nil, err)
		} else {
			output(user, nil)
		}
	})
}

func (s *UserRepo) AllUsers(output func(users []domain.User, err error)) {
	s.client.Records(func (users []domain.User, err error){
		if err != nil {
			output(nil, err)
		} else {
			output(users, nil)
		}

	})
}


