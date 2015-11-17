package main

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase"
	"net/http"
	"fmt"
	"github.com/byrnedo/apibase/db/mongo"
	"github.com/byrnedo/apibase/env"
)

func main() {

	var (
		host string
		port int
		err error
	)

	apibase.Init()

	mongo.Init(env.GetOr("MONGO_URL",apibase.Conf.GetDefaultString("mongo.url", "")), Trace)

	http.HandleFunc("/api/v1/healthcheck", healthCheck)


	host = apibase.Conf.GetDefaultString("http.host", "localhost")
	if port,err = env.GetOrInt("PORT", apibase.Conf.GetDefaultInt("http.port", 9999)); err != nil {
		panic(err.Error())
	}

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

