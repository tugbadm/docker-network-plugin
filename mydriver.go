package mydriver

import (
	"fmt"

	"github.com/docker/go-plugins-helpers/network"
)

type MyDriver struct {
	CreateNetworkBody network.CreateNetworkRequest
}

func NewDriver() *MyDriver {
	return new(MyDriver)
}

func (d *MyDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{
		Scope: LocalScope,
	}, nil
}

func (d *MyDriver) CreateNetwork(request *network.CreateNetworkRequest) error {
	/*	interfaces, err := net.Interfaces()
		if err != nil {
			return &network.ErrorResponse{
				Err: "cannot get network interfaces",
			}
		}
			for _, i := range interfaces {
				if i.Name == request.Options.Generic.Bridge {
					return nil
				}
			}

			return &network.ErrorResponse{
				Err: fmt.Sprintf(`network interface "%s" does not exist`, request.Options.Generic.Bridge),
			}
	*/
	return nil
}

func (d *MyDriver) DeleteNetwork(request *network.DeleteNetworkRequest) error {

	return nil
}

func (d *MyDriver) CreateEndpoint(request *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	// todo: create veth
	fmt.Printf("Request:\n%+v\n\n", request)

	/*veth := &netlink.Veth{}
	veth.LinkAttrs.Name = "veth1"
	err := netlink.LinkAdd(veth)

	if err != nil {
		c.Fail(http.StatusNotAcceptable, fmt.Sprintf("cannot create veth for network %s: %+v", request.NetworkID, err))
		return nil
	}*/

	return &network.CreateEndpointResponse{}, nil
}

func (d *MyDriver) DeleteEndpoint(request *network.DeleteEndpointRequest) error {

	return nil
}

func (d *MyDriver) EndpointInfo(request *network.InfoRequest) (*network.InfoResponse, error) {

	return nil, nil
}

func (d *MyDriver) Join(request *network.JoinRequest) (*network.JoinResponse, error) {

	return nil, nil
}

func (d *MyDriver) Leave(request *network.LeaveRequest) error {

	return nil
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
