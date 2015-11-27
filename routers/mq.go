package routers

import (
	"github.com/byrnedo/apibase/natsio"
	"github.com/byrnedo/blogsvc/controllers/mq"
	"github.com/byrnedo/apibase/controllers"
	. "github.com/byrnedo/apibase/logger"
)

func InitMq(url string) {

	var natsOpts *natsio.NatsOptions
	var natsCon *natsio.Nats = &natsio.Nats{}
	natsOpts = natsio.NewNatsOptions(func(n *natsio.NatsOptions) error {
		n.Url = url
		return nil
	})


	natsCon, err := natsOpts.ConnectOrRetry(3)
	if err != nil {
		panic("Failed to connect to nats:" + err.Error())
	}

	SubscribeRoutes(natsCon, mq.NewHealthcheckController(natsCon.EncCon))
	SubscribeRoutes(natsCon, mq.NewPostsController(natsCon.EncCon))
}

func SubscribeRoutes(natsCon *natsio.Nats, controllers controllers.NatsController) {
	for _, route := range controllers.GetRoutes() {
		Info.Printf("Subscribing handler for route %s\n", route.GetPath())
		natsCon.QueueSubscribe(route.GetPath(), "blog_svc_worker", route.GetHandler())
	}
}
