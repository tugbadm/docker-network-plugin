package mydriver

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

type VoidResponse struct{}

// plugin activate
type pluginActivateResponse struct {
	Implements []string
}

func pluginActivate() pluginActivateResponse {
	return pluginActivateResponse{
		Implements: []string{"NetworkDriver"},
	}
}
