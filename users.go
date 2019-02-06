package hallpass

import "errors"

// User represents a user in the system
type User struct {
	ID        int64  `json:"id,omitempty" form:"id,omitempty" db:"user_id, primarykey, autoincrement"`
	FirstName string `json:"first_name" form:"first_name" db:"first_name,size:64"`
	LastName  string `json:"last_name" form:"last_name" db:"last_name,size:64"`
	Email     string `json:"email" form:"email" db:"email,size:128"`
	Org       string `json:"org" form:"org" db:"org,size:128"`
	Passwd    string `json:"passwd,omitempty" form:"passwd,omitempty" db:"passwd"`
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
	Email string `url:"email,omitempty"`
	Org   string `url:"org,omitempty"`
}

// ErrUserNotFound is an error indicating a specified user could not be found
var ErrUserNotFound = errors.New("user not found")
