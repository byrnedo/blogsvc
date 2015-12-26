package daos

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/blogsvc/models"
	"github.com/byrnedo/blogsvc/msgspec/webmsgspec"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "posts"
)

type PostDAO interface {
	Find(bson.ObjectId) (*models.PostModel, error)
	FindMany(query map[string]interface{}, sortBy []string, offset int, limit int) ([]*models.PostModel, error)
	Create(*webmsgspec.NewPostDTO) (*models.PostModel, error)
	Replace(*webmsgspec.UpdatePostDTO) (*models.PostModel, error)
	Delete(bson.ObjectId) error
}

type DefaultPostDAO struct {
	Session *mgo.Session
}

func init() {
}

func NewDefaulPostDAO(session *mgo.Session) *DefaultPostDAO {
	dao := &DefaultPostDAO{session}
	dao.Ensures()
	return dao
}

func (u *DefaultPostDAO) Ensures() {
	index := mgo.Index{
		Key:        []string{"slug"},
		Unique:     true,
		DropDups:   false,
		Background: false, // See notes.
		Sparse:     true,
	}
	if err := u.col().EnsureIndex(index); err != nil {
		panic("Failed to create index:" + err.Error())
	}
}

func (uM *DefaultPostDAO) col() *mgo.Collection {
	return uM.Session.DB("").C(collection)
}

func (uM *DefaultPostDAO) Find(id bson.ObjectId) (u *models.PostModel, err error) {
	u = &models.PostModel{}
	q := uM.col().FindId(id).One(u)
	return u, q
}

func (uM *DefaultPostDAO) FindByEmail(email string) (u *models.PostModel, err error) {
	u = &models.PostModel{}
	q := uM.col().Find(bson.M{"email": email}).One(u)
	return u, q
}

func (uM *DefaultPostDAO) Create(nPost *webmsgspec.NewPostDTO) (u *models.PostModel, err error) {
	if u, err = nPost.MapToEntity(); err != nil {
		return
	}

	return u, uM.col().Insert(u)
}

func (uM *DefaultPostDAO) Replace(updPost *webmsgspec.UpdatePostDTO) (u *models.PostModel, err error) {
	if u, err = updPost.MapToEntity(); err != nil {
		return
	}
	var id = u.ID
	u.ID = ""

	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err = uM.col().Find(bson.M{"_id": id}).Apply(change, u)
	return
}

func (uM *DefaultPostDAO) Delete(id bson.ObjectId) error {
	return uM.col().RemoveId(id)
}

func (uM *DefaultPostDAO) FindMany(query map[string]interface{}, sortBy []string, offset int, limit int) ([]*models.PostModel, error) {
	var (
		err    error
		result = make([]*models.PostModel, 0)
	)

	err = uM.col().Find(query).Skip(offset).Limit(limit).Sort(sortBy...).All(&result)
	return result, err
}
