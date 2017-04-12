package domain

import "errors"

var (
	// ErrUserNotFound is when the user does not exist.
	ErrUserNotFound = errors.New("User not found.")
	// ErrUserAlreadyExist is when the user already exists.
	ErrUserAlreadyExist = errors.New("User already exists.")
	// ErrUserPasswordNotMatch is when the user's password does not match.
	ErrUserPasswordNotMatch = errors.New("User password does not match.")
)

// User represents a user of the system.
type User struct { //FIXME: should not have json tags in here
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// UserRepo represents a service for storage of users.
type UserRepo interface {
	FindByEmail(email string) (*User, error)
	Store(item *User) error
}
