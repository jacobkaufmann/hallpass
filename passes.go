package hallpass

import (
	"errors"
	"time"
)

// Pass represents a pass created by a user in the system
type Pass struct {
	ID             int       `json:"id,omitempty" sql:"id,omitempty"`
	Creator        User      `json:"creator"`
	Student        string    `json:"student" sql:"student"`
	CreatedDate    time.Time `json:"created" sql:"created_date"`
	ExpirationDate time.Time `json:"expiration" sql:"expiration_date"`
}

// PassesService interacts with the pass-related endponits in the API
type PassesService interface {
	// Get a pass
	Get(id int) (*Pass, error)

	// List passes
	List(opt *PassListOptions) ([]*Pass, error)

	// Create a pass
	Create(p *Pass) (bool, error)

	// Update a pass
	Update(p *Pass) (bool, error)
}

// PassListOptions provide options for retrieving a list of passes
type PassListOptions struct {
	CreatorEmail  string
	Student       string
	ExpiredBefore time.Time
	ExpiringAfter time.Time
}

// ErrPassNotFound is an error indicating a specified pass could not be found
var ErrPassNotFound = errors.New("pass not found")
