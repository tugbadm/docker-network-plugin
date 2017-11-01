package main

import (
	"github.com/docker/go-plugins-helpers/network"
	"github.com/tugbadartici/docker-network-plugin"
	"log"
)

func main() {
	d := mydriver.NewDriver()
	h := network.NewHandler(d)
	err := h.ServeTCP("test", ":8010", "", nil)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}
