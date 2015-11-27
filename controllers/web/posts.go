package web

import (
	routes "github.com/byrnedo/apibase/routes"
	"net/http"
)

type PostsController struct {
	routes []*routes.WebRoute
}

func (pC *PostsController) GetRoutes() []*routes.WebRoute{
	return pC.routes
}

func NewPostsController() (pC *PostsController){
	pC = &PostsController{}
	pC.routes = []*routes.WebRoute{
		routes.NewWebRoute("GetPosts", "/api/v1/posts", routes.GET, pC.List),
		routes.NewWebRoute("NewPost", "/api/v1/posts/{id}", routes.POST, pC.List),
		routes.NewWebRoute("UpdatePost", "/api/v1/posts/{id}", routes.PUT, pC.List),
		routes.NewWebRoute("DeletePost", "/api/v1/posts/{id}", routes.DELETE, pC.List),
	}
	return
}

func (pC *PostsController) List(http.ResponseWriter, *http.Request){

}
