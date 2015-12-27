package webcontrollers

import (
	routes "github.com/byrnedo/apibase/routes"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
"github.com/byrnedo/blogsvc/msgspec/webmsgspec"
	"github.com/byrnedo/svccommon/validate"
"gopkg.in/mgo.v2/bson"
	"github.com/byrnedo/apibase/controllers"
"github.com/byrnedo/blogsvc/daos"
	"github.com/byrnedo/apibase/db/mongo/defaultmongo"
	. "github.com/byrnedo/apibase/logger"
	svcSpec "github.com/byrnedo/svccommon/msgspec/web"

"github.com/byrnedo/blogsvc/models"
)

type PostsController struct {
	*controllers.JsonController
	postModel daos.PostDAO
}

func NewPostsController() *PostsController {
	return &PostsController{
		JsonController: &controllers.JsonController{},
		postModel:      daos.NewDefaulPostDAO(defaultmongo.Conn()),
	}
}

func (pC *PostsController) GetRoutes() []*routes.WebRoute {
	return []*routes.WebRoute{
		routes.NewWebRoute("NewPost", "/v1/posts/:postId", routes.POST, pC.Create),
		routes.NewWebRoute("ReplacePost", "/v1/posts/:postId", routes.PUT, pC.Replace),
		routes.NewWebRoute("GetPost", "/v1/posts/:postId", routes.GET, pC.GetOne),
		routes.NewWebRoute("GetPosts", "/v1/posts", routes.GET, pC.List),
		routes.NewWebRoute("DeletePost", "/v1/posts/:postId", routes.DELETE, pC.Delete),
	}
}

func (pC *PostsController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var u webmsgspec.NewPostResource

	if err := decoder.Decode(&u); err != nil {
		Error.Println(err)
		panic("Failed to decode json:" + err.Error())
	}

	if valErrs := validate.ValidateStruct(u); len(valErrs) != 0 {
		errResponse := svcSpec.NewValidationErrorResonse(valErrs)
		pC.ServeWithStatus(w, errResponse, 400)
		return
	}

	inserted, err := pC.postModel.Create(u.Data)
	if err != nil {
		Error.Println("Error creating post:" + err.Error())
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(500), 500)
		return
	}
	pC.ServeWithStatus(w, inserted, 201)
}

func (pC *PostsController) Replace(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("postId")
	if !bson.IsObjectIdHex(id) {
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(404), 404)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var u webmsgspec.UpdatedPostResource

	if err := decoder.Decode(&u); err != nil {
		Error.Println(err)
		panic("Failed to decode json:" + err.Error())
	}

	u.Data.ID = id

	if valErrs := validate.ValidateStruct(u); len(valErrs) != 0 {
		errResponse := svcSpec.NewValidationErrorResonse(valErrs)
		pC.ServeWithStatus(w, errResponse, 400)
		return
	}

	inserted, err := pC.postModel.Replace(u.Data)
	if err != nil {
		Error.Println("Error updating post:" + err.Error())
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(500), 500)
		return
	}
	pC.ServeWithStatus(w, inserted, 200)
}

func (pC *PostsController) GetOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		id    string
		err   error
		objId bson.ObjectId
		post  *models.PostModel
	)

	id = ps.ByName("postId")
	if !bson.IsObjectIdHex(id) {
		Error.Println("Id is not object id")
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(404), 404)
		return
	}

	objId = bson.ObjectIdHex(id)

	if post, err = pC.postModel.Find(objId); err != nil {
		Error.Println("Failed to find post:" + err.Error())
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(404), 404)
		return
	}

	pC.Serve(w, &webmsgspec.PostResource{post})
}

func (pC *PostsController) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query := pC.QueryInterfaceMap(r, "query", &models.PostModel{})
	order, _ := r.URL.Query()["order"]
	offset, _ := pC.QueryInt(r, "offset")
	limit, _ := pC.QueryInt(r, "limit")

	posts, err := pC.postModel.FindMany(query, order, offset, limit)
	if err != nil {
		Error.Println("Failed to find posts:", err)
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(404), 404)
		return
	}
	pC.Serve(w, &webmsgspec.PostsResource{posts})
}

func (pC *PostsController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("postId")

	if !bson.IsObjectIdHex(id) {
		Error.Println("Not an object id:", id)
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(404), 404)
		return
	}

	if err := pC.postModel.Delete(bson.ObjectIdHex(id)); err != nil {
		Error.Println("Error deleting:", err)
		pC.ServeWithStatus(w, svcSpec.NewErrorResponse().AddCodeError(404), 404)
		return
	}

	pC.Serve(w, &webmsgspec.PostsResource{nil})
}

