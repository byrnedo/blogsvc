package mqcontrollers

import (
	"github.com/apcera/nats"
	r "github.com/byrnedo/apibase/routes"
	"github.com/byrnedo/apibase/natsio"
)

type HealthcheckController struct {
	routes []*r.NatsRoute
	natsCon *natsio.Nats
}

func (c *HealthcheckController) GetRoutes() []*r.NatsRoute {
	return c.routes
}

func NewHealthcheckController(nc *natsio.Nats) (hc *HealthcheckController) {
	hc = &HealthcheckController{}
	hc.natsCon = nc
	hc.routes = []*r.NatsRoute{
		r.NewNatsRoute("blog.healthcheck", hc.Healthcheck),
	}
	return
}

func (c *HealthcheckController) Healthcheck(m *nats.Msg) {
	c.natsCon.EncCon.Publish(m.Reply, "up up up")
}
