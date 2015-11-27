package msgspec

import (
	"time"
	"gopkg.in/bluesuncorp/validator.v8"
)

type NewPost struct {
	Post
	ID string `validate:"required"`
}

type Post struct {
	ID string
	Title string `validate:"required"`
	Slug string
	Body string
	AuthorID string
	PublishTime time.Time
	CreationTime time.Time
	UpdateTime time.Time
}

func (p *NewPost) Validate() map[string]*validator.FieldError {
	return ValidateStruct(p)
}
func (p *Post) Validate() map[string]*validator.FieldError {
	return ValidateStruct(p)
}

