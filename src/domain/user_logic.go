package domain

// UserInteractor represents a service for managing users.
type Userinteractor struct {
	userRepo UserRepo
	passhash PasshashCase
}

// UserRepo represents a service for storage of users.
type UserRepo interface {
	FindByEmail(email string) (*User, error)
	Store(item *User) error
}

// NewUserCase returns the service for managing users.
func NewUserInteractor(repo domain.UserRepo, passhash domain.PasshashCase) *UserInteractor {
	s := new(UserInteractor)
	s.userRepo = repo
	s.passhash = passhash
	return s
}

// Authenticate returns an error if the email and password don't match.
func (s *UserInteractor) Authenticate(item *domain.User) error {
	q, err := s.userRepo.FindByEmail(item.Email)
	if err != nil {
		return domain.ErrUserNotFound
	}

	// If passwords match.
	if s.passhash.Match(q.Password, item.Password) {
		return nil
	}

	return domain.ErrUserPasswordNotMatch
}

// User returns a user by email.
func (s *UserInteractor) User(email string) (*domain.User, error) {
	item, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return item, domain.ErrUserNotFound
	}

	return item, nil
}

// CreateUser creates a new user.
func (s *UserInteractor) CreateUser(item *domain.User) error {
	_, err := s.userRepo.FindByEmail(item.Email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	}

	passNew, err := s.passhash.Hash(item.Password)
	if err != nil {
		return domain.ErrPasswordHash
	}

	// Swap the password.
	passOld := item.Password
	item.Password = passNew
	err = s.userRepo.Store(item)

	// Restore the password.
	item.Password = passOld

	return err
}
