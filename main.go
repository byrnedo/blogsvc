package main

import (
	. "github.com/byrnedo/apibase/logger"
	"github.com/byrnedo/apibase"
	"net/http"
	"fmt"
)

func main() {

	var (
		host string
		port int
	)

	apibase.Init()

	http.HandleFunc("/api/v1/healthcheck", healthCheck)

	host = apibase.Conf.GetDefaultString("http.host", "localhost")
	port = apibase.Conf.GetDefaultInt("http.port", 9999)

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	http.ListenAndServe(listenAddr, nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

