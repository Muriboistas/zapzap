package entity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User data
type User struct {
	ID        ID
	Name      string
	URN       string
	Language  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser create a new user
func NewUser(name, urn, language string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Name:      name,
		URN:       urn,
		Language:  language,
		CreatedAt: time.Now(),
	}

	err := u.Validate()
	if err != nil {
		return nil, err
	}

	return u, err
}

// Validate our user
func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(1, 25)),
		validation.Field(&u.URN, validation.Required),
		validation.Field(&u.Language, validation.Required),
	)
}
