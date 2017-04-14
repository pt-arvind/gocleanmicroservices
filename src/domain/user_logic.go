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
