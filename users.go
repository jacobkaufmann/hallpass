package hallpass

import "errors"

// User represents a user in the system
type User struct {
	ID        int64  `json:"id,omitempty" form:"id,omitempty" sql:"user_id"`
	FirstName string `json:"first_name" form:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" form:"last_name" sql:"last_name"`
	Email     string `json:"email" form:"email" sql:"email"`
	Org       string `json:"org" form:"org" sql:"org"`
	Passwd    string `json:"passwd,omitempty" form:"passwd,omitempty" sql:"passwd"`
}

// UsersService interacts with the user-related endpoints in the API
type UsersService interface {
	// Get a user
	Get(id int64) (*User, error)

	// List users
	List(opt *UserListOptions) ([]*User, error)

	// Create a user
	Create(u *User) (bool, error)

	// Delete a user
	Delete(id int64) (bool, error)

	// Update a user
	Update(u *User) (bool, error)
}

// UserListOptions provide options for retrieving a list of users
type UserListOptions struct {
	Email string
	Org   string
}

// ErrUserNotFound is an error indicating a specified user could not be found
var ErrUserNotFound = errors.New("user not found")
