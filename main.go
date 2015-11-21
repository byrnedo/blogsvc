package main

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase"
	"net/http"
	"fmt"
	"github.com/byrnedo/apibase/db/mongo"
	"github.com/byrnedo/apibase/env"
	"github.com/byrnedo/apibase/natsio"
	"github.com/apcera/nats"
)

func main() {

	var (
		host string
		port int
		err error
	)

	apibase.Init()

	setupMongo()

	setupNats()

	http.HandleFunc("/api/v1/healthcheck", healthCheck)

	host = apibase.Conf.GetDefaultString("http.host", "localhost")
	if port,err = env.GetOrInt("PORT", apibase.Conf.GetDefaultInt("http.port", 9999)); err != nil {
		panic(err.Error())
	}

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func setupMongo(){
	mongo.Init(env.GetOr("MONGO_URL",apibase.Conf.GetDefaultString("mongo.url", "")), Trace)
}

func setupNats(){
	var natsOpts *natsio.Nats
	natsOpts = natsio.NewNats(func(n *natsio.Nats) error{
		n.Url = env.GetOr("NATS_URL", apibase.Conf.GetDefaultString("nats.url", "nats://localhost:4222"))
		return nil
	})
	natsOpts.HandleFunc("blog.healthcheck", func(m *nats.Msg){

	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

