package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PostModel struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Body         string        `json:"body"`
	AuthorID     string        `json:"author"`
	PublishTime  time.Time     `json:"publish_time"`
	CreationTime time.Time     `bson:"creationtime,omitempty" json:"creation_time"`
	UpdateTime   time.Time     `json:"update_time"`
}
