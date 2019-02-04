package hallpass

import "time"

// Pass represents a pass created by a user in the system
type Pass struct {
	ID             int       `json:"id,omitempty"`
	Giver          User      `json:"giver"`
	Student        string    `json:"student"`
	TimeGiven      time.Time `json:"time_given"`
	TimeExpiration time.Time `json:"time_expiration"`
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
	GiverEmail        string
	Student           string
	ExpiredBeforeTime time.Time
	ExpiringAfterTime time.Time
}
