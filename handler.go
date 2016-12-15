package mydriver

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/vishvananda/netlink"
	go_docker "github.com/fsouza/go-dockerclient"
)

type VoidResponse struct{}

type Docker struct{
	client *go_docker.Client
}


func generateVethName() string {
	n := 6
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return fmt.Sprintf("veth-%s", string(b))
}

func (d *MyDriver) getLinkByName(name string) (netlink.Link, error) {
	l, err := netlink.LinkByName(name)
	if err != nil{
		return nil, err
	}
	if l == nil {
		return nil, fmt.Errorf(`link "%s" does not exist`, name)
	}
	return l, nil
}

func (d *MyDriver) deleteLinkByName(name string) error{
	l, err := d.getLinkByName(name)
	if err != nil {
		return err
	}

	err = netlink.LinkSetDown(l)
	if err != nil {
		return fmt.Errorf("cannot set link %s down", name)
	}

	err = netlink.LinkDel(l)
	if err != nil {
		return fmt.Errorf("cannot delete link %s ", name)
	}
	return nil
}

func (d *MyDriver) LinkSetUpByName(name string) error{
	l, err := d.getLinkByName(name)
	if err != nil{
		return err
	}

	return netlink.LinkSetUp(l)
}

// GetVMContainer returns the underlying docker container object for a
// given VM
func (d *Docker) GetContainerInfo (endpointID string) (*go_docker.Container, error) {
	containers, err := d.client.ListContainers(go_docker.ListContainersOptions{
		Filters: map[string][]string{
			"name": {
				endpointID,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	// no such container
	if len(containers) == 0 {
		return nil, nil
	}

	return d.client.InspectContainer(containers[0].ID)
}