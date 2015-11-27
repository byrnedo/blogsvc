package mq

import (
	r "github.com/byrnedo/apibase/routes"
	"github.com/apcera/nats"
)

type PostsController struct {
	routes []*r.NatsRoute
	encCon *nats.EncodedConn
}

func (c *PostsController) GetRoutes() []*r.NatsRoute {
	return c.routes
}

func NewPostsController(nc *nats.EncodedConn) (pC *PostsController) {
	pC = &PostsController{}
	pC.encCon = nc
	pC.routes = []*r.NatsRoute{
		r.NewNatsRoute("blog.posts.list", pC.List),
	}
	return
}

func (c *PostsController) List(m *nats.Msg) {
	c.encCon.Publish(m.Reply, "Not implemented")
}
