package routers

import (
	"github.com/byrnedo/apibase/controllers"
	"github.com/byrnedo/blogsvc/controllers/mqcontrollers"
	"github.com/byrnedo/apibase/natsio/defaultnats"
)

func init() {
	controllers.SubscribeNatsRoutes(defaultnats.Conn, "post_svc_worker", mqcontrollers.NewHealthcheckController(defaultnats.Conn))
	controllers.SubscribeNatsRoutes(defaultnats.Conn, "post_svc_worker", mqcontrollers.NewPostsController(defaultnats.Conn))
}
