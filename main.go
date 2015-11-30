package main

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase"
	"net/http"
	"fmt"
	"github.com/byrnedo/apibase/env"
	"github.com/byrnedo/blogsvc/routers"
	"github.com/byrnedo/apibase/db/postgres"
)

func main() {

	var (
		host string
		port int
		err error
	)

	apibase.Init()

	//mongo.Init(env.GetOr("MONGO_URL", apibase.Conf.GetDefaultString("mongo.url", "")), Trace)
	postgres.Init(func(c *postgres.Config){
		var psqlCon string
		if psqlCon, err = apibase.Conf.GetString("postgres.connect-string"); err != nil {
			panic("Failed to connect to postgres:" + err.Error())
		}
		c.ConnectString = env.GetOr("POSTGRES_CON_STRING", psqlCon )
	})

	postgres.Migrate(apibase.Conf.GetDefaultString("postgres.migrations", "./conf/migrations"))

	routers.InitMq(env.GetOr("NATS_URL", apibase.Conf.GetDefaultString("nats.url", "nats://localhost:4222")))

	webRouter := routers.InitWeb()

	host = apibase.Conf.GetDefaultString("http.host", "localhost")
	if port, err = env.GetOrInt("PORT", apibase.Conf.GetDefaultInt("http.port", 9999)); err != nil {
		panic(err.Error())
	}

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	http.ListenAndServe(listenAddr, webRouter)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

