package mydriver

import (
	"gongular"
	"log"

	"github.com/docker/go-plugins-helpers/network"
)

const (
	manifest = `{"Implements": ["NetworkDriver"]}`
	// LocalScope is the correct scope response for a local scope driver
	LocalScope = `local`
	// GlobalScope is the correct scope response for a global scope driver
	GlobalScope        = `global`
	pluginActivatePath = "/Plugin.Activate"
	capabilitiesPath   = "/NetworkDriver.GetCapabilities"
	createNetworkPath  = "/NetworkDriver.CreateNetwork"
	deleteNetworkPath  = "/NetworkDriver.DeleteNetwork"
	createEndpointPath = "/NetworkDriver.CreateEndpoint"
	endpointInfoPath   = "/NetworkDriver.EndpointOperInfo"
	deleteEndpointPath = "/NetworkDriver.DeleteEndpoint"
	joinPath           = "/NetworkDriver.Join"
	leavePath          = "/NetworkDriver.Leave"
	discoverNewPath    = "/NetworkDriver.DiscoverNew"
	discoverDeletePath = "/NetworkDriver.DiscoverDelete"
	programExtConnPath = "/NetworkDriver.ProgramExternalConnectivity"
	revokeExtConnPath  = "/NetworkDriver.RevokeExternalConnectivity"
)

// gongular accepts "xxBody" syntax for POST requests
type (
	DeleteNetworkBody               network.DeleteNetworkRequest
	CreateEndpointBody              network.CreateEndpointRequest
	DeleteEndpointBody              network.DeleteEndpointRequest
	InfoBody                        network.InfoRequest
	JoinBody                        network.JoinRequest
	LeaveBody                       network.LeaveRequest
	DiscoveryNotificationBody       network.DiscoveryNotification
	ProgramExternalConnectivityBody network.ProgramExternalConnectivityRequest
	RevokeExternalConnectivityBody  network.RevokeExternalConnectivityRequest
)

type VoidResponse struct{}

// plugin activate
type pluginActivateResponse struct {
	Implements []string
}

type CreateNetworkBody struct {
	network.CreateNetworkRequest
	Options CreateNetworkOptions
}

// Create network options
type CreateNetworkGeneric struct {
	Bridge string `json:"bridge"`
}

type CreateNetworkOptions struct {
	EnableIPv6 bool                 `json:"com.docker.network.enable_ipv6"`
	Generic    CreateNetworkGeneric `json:"com.docker.network.generic"`
}

func (d *MyDriver) Handle(port string) {
	r := gongular.NewRouter()
	r.POST(pluginActivatePath, pluginActivate)
	r.POST(capabilitiesPath, d.GetCapabilities)
	r.POST(createNetworkPath, d.CreateNetwork)
	r.POST(deleteNetworkPath, d.DeleteNetwork)
	r.POST(createEndpointPath, d.CreateEndpoint)
	r.POST(joinPath, d.Join)

	err := r.ListenAndServe(port)
	if err != nil {
		log.Fatal(err)
	}
}

func pluginActivate(c *gongular.Context) pluginActivateResponse {
	return pluginActivateResponse{
		Implements: []string{"NetworkDriver"},
	}
}
