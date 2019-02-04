package hallpass

// User represents a user in the system
type User struct {
	ID        int    `json:"id,omitempty" form:"id,omitempty"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Passwd    string `json:"passwd,omitempty" form:"passwd,omitempty"`
}

// UsersService interacts with the user-related endpoints in the API
type UsersService interface {
	// Get a user
	Get(id int) *User

	// Create a user
	Create(u *User) bool

	// Delete a user
	Delete(id int) *User

	// Update a user
	Update(u *User) bool
}
