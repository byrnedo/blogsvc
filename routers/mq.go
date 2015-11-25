package routers

import (
	"github.com/byrnedo/apibase/natsio"
	"github.com/byrnedo/blogsvc/controllers/mq"
	"github.com/byrnedo/apibase/controllers"
)

func InitMq(url string){
	var natsOpts *natsio.NatsOptions
	var natsCon *natsio.Nats
	natsOpts = natsio.NewNatsOptions(func(n *natsio.NatsOptions) error{
		n.Url = url
		return nil
	})

	con1 := mq.NewHealthcheckController(natsCon.EncCon)

	bindRoutes(natsOpts, con1)

	natsCon, err := natsOpts.ListenAndServe()
	if err != nil {
		panic("Failed to connect to nats:" + err.Error())
	}
}

func bindRoutes(opts *natsio.NatsOptions, con controllers.NatsController){
	for _, route := range con.GetRoutes() {
		opts.HandleFunc(route.GetPath(), route.GetHandler())
	}
}
