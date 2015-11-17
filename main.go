package main

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase"
	"net/http"
	"fmt"
	"github.com/byrnedo/apibase/db/mongo"
	"os"
	"strconv"
)

func GetEnvOr(key string, fallback string) string {
	if val, found := os.LookupEnv(key); found {
		return val
	}
	return fallback
}

func GetEnvOrInt(key string, fallback int) int {
	if strVal, found := os.LookupEnv(key); found != false {
		if val,err := strconv.Atoi(strVal); err != nil {
			return val
		} else {
			panic("Failed to make int from ENV " + key + ": " + err.Error())
		}
	} else {
		return fallback
	}
}



func main() {

	var (
		host string
		port int
	)

	apibase.Init()

	mongo.Init(GetEnvOr("MONGO_URL",apibase.Conf.GetDefaultString("mongo.url", "")), Trace)

	http.HandleFunc("/api/v1/healthcheck", healthCheck)


	host = apibase.Conf.GetDefaultString("http.host", "localhost")
	port = GetEnvOrInt("PORT", apibase.Conf.GetDefaultInt("http.port", 9999))

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

