package routers

import (
	"github.com/byrnedo/apibase/controllers"
	"github.com/byrnedo/apibase/natsio"
	"github.com/byrnedo/apibase/natsio/defaultnats"
	"github.com/byrnedo/blogsvc/controllers/mqcontrollers"
	"time"
)

func init() {
	controllers.SubscribeNatsRoutes(defaultnats.Conn, "post_svc_worker", mqcontrollers.NewHealthcheckController(defaultnats.Conn))
	controllers.SubscribeNatsRoutes(defaultnats.Conn, "post_svc_worker", mqcontrollers.NewPostsController(defaultnats.Conn))
}
