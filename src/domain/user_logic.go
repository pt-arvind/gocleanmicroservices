package domain

// UserInteractor represents a service for managing users.
type Userinteractor struct {
	userRepo UserRepo
	passhash PasshashCase
}

// UserRepo represents a service for storage of users.
type UserRepo interface {
	FindByEmail(email string, output func(user *User, err error))
	Store(item User, output func(user *User, err error))
	AllUsers(output func(users []User, err error))

}
