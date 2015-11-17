package main

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase"
	"net/http"
	"fmt"
	"github.com/byrnedo/apibase/db/mongo"
	"os"
)

func GetEnvOr(key string, fallback string) string {
	if val, found := os.LookupEnv; found {
		return val
	}
	return fallback
}


func main() {

	var (
		host string
		port int
	)

	apibase.Init()

	mongo.Init(apibase.Conf.GetDefaultString("mongo.url", GetEnvOr("MONGO_URL", "")), Trace)

	http.HandleFunc("/api/v1/healthcheck", healthCheck)

	host = apibase.Conf.GetDefaultString("http.host", "localhost")
	port = apibase.Conf.GetDefaultInt("http.port", GetEnvOr("PORT", 9999))

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

