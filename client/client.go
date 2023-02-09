package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	runtime2 "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"

	"github.com/hashicorp/waypoint-client/config"
	namespace "github.com/hashicorp/waypoint-client/context"
	smwaypoint "github.com/hashicorp/waypoint-client/gen/client/waypoint"
	wcs "github.com/hashicorp/waypoint-client/gen/client/waypoint_control_service"
)

const (
	defaultBasePath = "/"
	defaultHCPHost  = "api.cloud.hashicorp.com"
)

type Config struct {
	//HCPConfig contains config values to interact with HCP Waypoint API
	*config.HCPWaypointConfig

	//WaypointConfig contains config values to interact with Self Managed Waypoint API.
	*config.WaypointConfig

	//BasePath is the path to the api server.
	BasePath string
}

func (c *Config) isHCP() bool {
	return c.HCPWaypointConfig != nil
}

func (c *Config) isSelfManaged() bool {
	return c.WaypointConfig != nil
}

// New creates a new Waypoint client given the config.
// If configured for HCP Waypoint Get Namespace and attach it to the basepath and setup bearer token transport.
func New(config Config) (smwaypoint.ClientService, error) {

	tlsTransport := cleanhttp.DefaultPooledTransport()
	tlsTransport.TLSClientConfig = &tls.Config{}

	var transport http.RoundTripper = &oauth2.Transport{
		Base:   tlsTransport,
		Source: &config,
	}
	if config.BasePath == "" {
		config.BasePath = "/"
	}

	var authWriter runtime2.ClientAuthInfoWriter
	if config.isHCP() {
		if config.ServerUrl == "" {
			config.ServerUrl = defaultHCPHost
		}

		token, err := config.HCPWaypointConfig.Token()
		if err != nil {
			return nil, fmt.Errorf("Cannot retrieve token from current HCP Credentials.")
		}

		//create client to get namespace only. this client is only used once on initialization to grab the namespace.
		runtime := httptransport.New(config.ServerUrl, config.BasePath, []string{"https"})

		authWriter := httptransport.BearerToken(token.AccessToken)
		runtime.DefaultAuthentication = authWriter

		namespaceClient := wcs.New(runtime, strfmt.Default)
		//Do GetNamespace Call. Put it in the context.
		namespaceId, err := namespace.GetNamespace(namespaceClient, config.HCPOrgId, config.HCPProjectId)
		if err != nil {
			return nil, err
		}
		if namespaceId == "" || namespaceId == "0" {
			return nil, fmt.Errorf("No namespace found associated with orgId and projectId provided.")
		}

		// Add namespace Id to the basepath
		basePath := config.BasePath + "waypoint/2022-02-03/namespace/" + namespaceId
		config.BasePath = basePath
	} else if config.isSelfManaged() {
		authWriter = httptransport.APIKeyAuth("authorization", "header", config.WaypointConfig.WaypointUserToken)
	}

	//Initialize API Client.
	runtime := httptransport.New(config.ServerUrl, config.BasePath, []string{"https"})
	runtime.Transport = transport
	runtime.DefaultAuthentication = authWriter

	apiClient := smwaypoint.New(runtime, strfmt.Default)

	return apiClient, nil
}
