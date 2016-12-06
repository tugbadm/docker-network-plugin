package mydriver

import (
	"fmt"
	"gongular"
	"net"
	"net/http"

	"github.com/docker/go-plugins-helpers/network"
)

type MyDriver struct {
	CreateNetworkBody network.CreateNetworkRequest
}

func NewDriver() *MyDriver {
	return new(MyDriver)
}

func (d *MyDriver) GetCapabilities(c *gongular.Context) (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{
		Scope: LocalScope,
	}, nil
}

func (d *MyDriver) CreateNetwork(c *gongular.Context, request CreateNetworkBody) *VoidResponse {
	interfaces, err := net.Interfaces()
	if err != nil {
		c.Fail(http.StatusNotAcceptable, network.ErrorResponse{
			Err: fmt.Sprintf("cannot reach network interfaces: %+v", err),
		})
		return nil
	}

	for _, i := range interfaces {
		if i.Name == request.Options.Generic.Bridge {
			return &VoidResponse{}
		}
	}

	c.Fail(http.StatusBadRequest, network.ErrorResponse{
		Err: fmt.Sprintf(`network interface "%s" does not exist`, request.Options.Generic.Bridge),
	})
	return nil
}

func (d *MyDriver) DeleteNetwork(c *gongular.Context, request *DeleteNetworkBody) error {

	return nil
}

func (d *MyDriver) CreateEndpoint(c *gongular.Context, request *CreateEndpointBody) (*network.CreateEndpointResponse, error) {

	return &network.CreateEndpointResponse{}, nil
}

func (d *MyDriver) DeleteEndpoint(c *gongular.Context, request *DeleteEndpointBody) error {

	return nil
}

func (d *MyDriver) EndpointInfo(c *gongular.Context, request *InfoBody) (*network.InfoResponse, error) {

	return nil, nil
}

func (d *MyDriver) Join(c *gongular.Context, request *JoinBody) (*network.JoinResponse, error) {

	return nil, nil
}

func (d *MyDriver) Leave(c *gongular.Context, request *LeaveBody) error {

	return nil
}

func (d *MyDriver) DiscoverNew(c *gongular.Context, request *DiscoveryNotificationBody) error {

	return nil
}

func (d *MyDriver) DiscoverDelete(c *gongular.Context, request *DiscoveryNotificationBody) error {

	return nil
}

func (d *MyDriver) ProgramExternalConnectivity(c *gongular.Context, request *ProgramExternalConnectivityBody) error {

	return nil
}

func (d *MyDriver) RevokeExternalConnectivity(c *gongular.Context, request *RevokeExternalConnectivityBody) error {

	return nil
}
