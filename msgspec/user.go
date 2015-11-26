package msgspec

import (
	"time"
	validator "gopkg.in/bluesuncorp/validator.v8"
)

type User struct {
	Alias string `validate:"required"`
	FirstName string
	LastName string
	Email string `validate:"required"`
	Password string
	CreationTime time.Time
	UpdateTime time.Time
}

func (u *User) Validate() map[string]*validator.FieldError {
	return V.Struct(u)
}

