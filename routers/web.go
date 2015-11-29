package routers

import (
	"github.com/gorilla/mux"
	"github.com/byrnedo/blogsvc/controllers/web"
	"github.com/byrnedo/apibase/controllers"
)

func InitWeb() (rtr *mux.Router) {
	rtr = mux.NewRouter().StrictSlash(true)
	controllers.RegisterMuxRoutes(rtr, &web.PostsController{})
	return
}

