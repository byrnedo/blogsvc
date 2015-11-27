package routers

import (
	"github.com/gorilla/mux"
	"github.com/byrnedo/blogsvc/controllers/web"
	"github.com/byrnedo/apibase/controllers"
)

func InitWeb() (rtr *mux.Router) {
	rtr = mux.NewRouter().StrictSlash(true)
	RegisterRoutes(rtr, web.NewPostsController())
	return
}

func RegisterRoutes(rtr *mux.Router, controller controllers.WebController){
	for _, route := range controller.GetRoutes() {
		rtr.
		Methods(route.GetMethod()).
		Path(route.GetPath()).
		Name(route.GetName()).
		Handler(route.GetHandler())
	}
}