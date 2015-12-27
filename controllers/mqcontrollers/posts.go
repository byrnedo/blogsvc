package mqcontrollers

import (
	"github.com/apcera/nats"
	r "github.com/byrnedo/apibase/routes"
	"github.com/byrnedo/apibase/natsio"
)

type PostsController struct {
	routes  []*r.NatsRoute
	natsCon *natsio.Nats
}

func (c *PostsController) GetRoutes() []*r.NatsRoute {
	return []*r.NatsRoute{
		r.NewNatsRoute("blog.posts.list", c.List),
	}
}

func NewPostsController(nc *natsio.Nats) (pC *PostsController) {
	pC = &PostsController{}
	pC.natsCon = nc
	return
}

func (c *PostsController) List(m *nats.Msg) {
	c.natsCon.EncCon.Publish(m.Reply, "Not implemented")
}
