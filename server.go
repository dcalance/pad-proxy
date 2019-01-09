package main

import (
	_ "net/http"

	"./config"
	"./handler"
	log "github.com/Sirupsen/logrus"
)

func main() {
	goproxy := handler.NewProxyServer()

	log.Infof("Start the proxy server in port:%s", config.RuntimeViper.GetString("server.port"))
	log.Fatal(goproxy.ListenAndServe())
}
