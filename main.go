package main

import (
	"fmt"
	"github.com/byrnedo/apibase/config"
	"github.com/byrnedo/apibase/helpers/envhelp"
	_ "github.com/byrnedo/apibase/natsio/defaultnats"
	_ "github.com/byrnedo/apibase/db/mongo/defaultmongo"
	. "github.com/byrnedo/apibase/logger"
	"net/http"
)

func main() {

	var (
		host string
		port int
		err  error
	)

	host = config.Conf.GetDefaultString("http.host", "localhost")
	if port, err = envhelp.GetOrInt("PORT", int(config.Conf.GetDefaultInt("http.port", 9999))); err != nil {
		panic(err.Error())
	}

	var listenAddr = fmt.Sprintf("%s:%d", host, port)
	Info.Printf("listening on " + listenAddr)
	if err = http.ListenAndServe(listenAddr, nil); err != nil {
		panic("Failed to start server:" + err.Error())
	}
}
