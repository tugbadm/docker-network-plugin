package main

import "github.com/docker-network-plugin"

func main() {
	d := mydriver.NewDriver()
	d.Handle(":8010")
}
