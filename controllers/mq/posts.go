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

func NewPostsController(nc *nats.EncodedConn) (hc *PostsController) {
	hc = &PostsController{}
	hc.encCon = nc
	hc.routes = []*r.NatsRoute{
		r.NewNatsRoute("blog.posts.list", hc.List),
	}
	return
}

func (c *PostsController) List(m *nats.Msg) {
	c.encCon.Publish(m.Reply, "Not implemented")
}
