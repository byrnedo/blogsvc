package msgspec

import (
	"time"
	"gopkg.in/bluesuncorp/validator.v8"
)

type Post struct {
	ID string
	Title string
	Slug string
	Body string
	AuthorID string
	PublishTime time.Time
	CreationTime time.Time
	UpdateTime time.Time
}

func (p *Post) Validate() map[string]*validator.FieldError {
	return V.Struct(p)
}
