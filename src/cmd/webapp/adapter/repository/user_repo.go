package repository

//adapter layer

import "domain"

// Service represents a service for interacting with the database.
type ServiceInput interface {
	Read() error
	Write() error
	Records() []domain.User
	AddRecord(domain.User)
}

type ServiceOutput interface {
	DidRead()
	DidStore()
	Error()
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

/*
	load
	get all users
	get user by email
	write
 */

// FindByEmail returns a user by an email.
func (s *UserRepo) FindByEmail(email string) (*domain.User, error) {
	item := new(domain.User)

	// Load the data.
	err := s.client.Read() //FIXME: shouldn't be here
	if err != nil {
		return item, err
	}

	// Determine if the record exists.
	for _, v := range s.client.Records() {
		if v.Email == email {
			return &v, nil
		}
	}

	return item, domain.ErrUserNotFound //TODO: should not be domain
}

// Store adds a user.
func (s *UserRepo) Store(item *domain.User) error {
	// Load the data.
	err := s.client.Read() //FIXME: shouldn't be here
	if err != nil {
		return err
	}

	// Add the record.
	s.client.AddRecord(*item)

	// Save the record to the database.
	return s.client.Write()
}



