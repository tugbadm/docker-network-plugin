package mydriver

import (
	"encoding/json"
	"fmt"
	"net"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/vishvananda/netlink"
	"github.com/fsouza/go-dockerclient"
)

type MyDriver struct {
	networks map[string]Network
}

type Network struct{
	bridge string
	endpoints map[string]Endpoint
}
type Endpoint struct{
	vethHost string
	vethCont string
}

func NewDriver() *MyDriver {
	d := new(MyDriver)
	return d
}

func (d *MyDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{
		Scope: network.LocalScope,
	}, nil
}

func (d *MyDriver) CreateNetwork(request *network.CreateNetworkRequest) error {
	options := request.Options["com.docker.network.generic"]
	bridge, ok := options.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data in request: %v", options)
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return fmt.Errorf("cannot get network interfaces")
	}

	for _, i := range interfaces {
		if i.Name == bridge["bridge"] { // bridge exists
			d.networks[request.NetworkID] = Network{
				bridge: i.Name,
				endpoints: map[string]Endpoint{},
			}
			return nil
		}
	}
	return fmt.Errorf(`bridge "%s" does not exist`, bridge["bridge"])
}

func (d *MyDriver) DeleteNetwork(request *network.DeleteNetworkRequest) error {
	return nil
}

func (d *MyDriver) CreateEndpoint(request *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	vethHost := generateVethName()
	vethCont := generateVethName()

	veth := &netlink.Veth{}
	veth.LinkAttrs.Name = vethCont
	veth.PeerName = vethHost

	err := netlink.LinkAdd(veth)
	if err != nil {
		return nil, fmt.Errorf(`cannot add link "%s"`, vethHost)
	}

	err = d.LinkSetUpByName(vethHost)
	if err != nil {
		return nil, err
	}

	err = d.LinkSetUpByName(vethCont)
	if err != nil {
		return nil, err
	}
	d.networks[request.NetworkID].endpoints[request.EndpointID] = Endpoint{
		vethCont: vethCont,
		vethHost: vethHost,
	}

	return &network.CreateEndpointResponse{}, nil

}

func (d *MyDriver) DeleteEndpoint(request *network.DeleteEndpointRequest) error {
	return nil
}

func (d *MyDriver) EndpointInfo(request *network.InfoRequest) (*network.InfoResponse, error) {
	endpointInfo := make(map[string]string)
	d := Docker{
		client: &docker.Client{},
	}

	endpoint, err := d.GetContainerInfo(request.EndpointID)
	if err != nil{
		return nil, err
	}

	out, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}
	endpointInfo[request.EndpointID] = string(out)
	return &network.InfoResponse{
		Value: endpointInfo,
	}, nil
}

func (d *MyDriver) Join(request *network.JoinRequest) (*network.JoinResponse, error) {
	vethHost, err := d.getLinkByName(d.networks[request.NetworkID].endpoints[request.EndpointID].vethHost)
	if err != nil {
		return nil, err
	}

	bridge, err := d.getLinkByName(d.networks[request.NetworkID].bridge)
	b, ok := bridge.(*netlink.Bridge)
	if !ok {
		return nil, fmt.Errorf(`bridge "%s" incorrect`, d.networks[request.NetworkID].bridge)
	}

	err = netlink.LinkSetMaster(vethHost, b)
	if err != nil {
		return nil, fmt.Errorf(`cannot add veth "%s" to bridge "%s"`, vethHost.Attrs().Name, d.networks[request.NetworkID].bridge)
	}

	return &network.JoinResponse{
		InterfaceName: network.InterfaceName{
			SrcName:  d.networks[request.NetworkID].endpoints[request.EndpointID].vethCont,
			DstPrefix: "eth",
		},
	}, nil
}

func (d *MyDriver) Leave(request *network.LeaveRequest) error {

	return d.deleteLinkByName(d.networks[request.NetworkID].endpoints[request.EndpointID].vethHost)
}

func (d *MyDriver) DiscoverNew(request *network.DiscoveryNotification) error {

	return nil
}

func (d *MyDriver) DiscoverDelete(request *network.DiscoveryNotification) error {
	return nil
}

func (d *MyDriver) ProgramExternalConnectivity(request *network.ProgramExternalConnectivityRequest) error {

	return nil
}

func (d *MyDriver) RevokeExternalConnectivity(request *network.RevokeExternalConnectivityRequest) error {

	return nil
}

func (d *MyDriver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {

	return nil, nil
}

func (d *MyDriver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}
