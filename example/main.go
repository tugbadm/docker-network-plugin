package main

import (
	"github.com/docker/go-plugins-helpers/network"
	"github.com/tugbadartici/docker-network-plugin"
)

func main() {
	d := mydriver.NewDriver()
	h := network.NewHandler(d)
	h.ServeTCP("test", ":8010", nil)
}
