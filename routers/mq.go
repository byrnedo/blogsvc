package routers

import (
	"github.com/byrnedo/apibase/natsio"
	"github.com/byrnedo/blogsvc/controllers/mq"
	"github.com/byrnedo/apibase/controllers"
	"time"
)

func InitMq(url string) {

	var natsOpts *natsio.NatsOptions
	var natsCon *natsio.Nats = &natsio.Nats{}
	natsOpts = natsio.NewNatsOptions(func(n *natsio.NatsOptions) error {
		n.Url = url
		n.Timeout = 10 * time.Second
		return nil
	})


	natsCon, err := natsOpts.Connect()
	if err != nil {
		panic("Failed to connect to nats:" + err.Error())
	}

	controllers.SubscribeNatsRoutes(natsCon, mq.NewHealthcheckController(natsCon.EncCon))
	controllers.SubscribeNatsRoutes(natsCon, mq.NewPostsController(natsCon.EncCon))
}

