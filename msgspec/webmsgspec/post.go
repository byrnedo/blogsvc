package webmsgspec

import (
	"github.com/byrnedo/blogsvc/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type NewPostResource struct {
	Data *NewPostDTO `json:"data" validate:"required"`
}

type UpdatedPostResource struct {
	Data *UpdatePostDTO `json:"data" validate:"required"`
}

type PostResource struct {
	Data *models.PostModel `json:"data"`
}

type PostsResource struct {
	Data []*models.PostModel `json:"data"`
}

type NewPostDTO struct {
	Title       string    `json:"title" validate:"required"`
	Slug        string    `json:"slug" validate:"required"`
	Body        string    `json:"body" validate:"omitempty"`
	AuthorID    string    `json:"author" validate:"required"`
	PublishTime time.Time `json:"publish_time"`
}

func (nU *NewPostDTO) MapToEntity() (*models.PostModel, error) {
	var (
		now = bson.Now()
	)

	return &models.PostModel{
		ID:           bson.NewObjectId(),
		Title:        nU.Title,
		Slug:         nU.Slug,
		Body:         nU.Body,
		AuthorID:     nU.AuthorID,
		CreationTime: now,
		UpdateTime:   now,
	}, nil
}

type UpdatePostDTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Slug        string    `json:"slug" validate:"required"`
	Body        string    `json:"body" validate:"omitempty"`
	AuthorID    string    `json:"author" validate:"required"`
	PublishTime time.Time `json:"publish_time"`
}

func (uU *UpdatePostDTO) MapToEntity() (*models.PostModel, error) {

	return &models.PostModel{
		ID:         bson.ObjectIdHex(uU.ID),
		Title:      uU.Title,
		Slug:       uU.Slug,
		Body:       uU.Body,
		AuthorID:   uU.AuthorID,
		UpdateTime: bson.Now(),
	}, nil
}
