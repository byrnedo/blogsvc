package routers

import (
	"github.com/byrnedo/apibase/controllers"
	"github.com/byrnedo/apibase/middleware"
	"github.com/byrnedo/blogsvc/controllers/webcontrollers"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func init() {
	var rtr = httprouter.New()
	controllers.RegisterRoutes(rtr, webcontrollers.NewPostsController())

	//alice is a tiny package to chain middlewares.
	handlerChain := alice.New(
		//limiterMw.Handler,
		middleware.LogTime,
		middleware.RecoverHandler,
		middleware.AcceptJsonHandler,
	).Then(rtr)

	http.Handle("/", handlerChain)
}
