package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	runtime2 "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"

	"github.com/hashicorp/waypoint-client/config"
	namespace "github.com/hashicorp/waypoint-client/context"
	smwaypoint "github.com/hashicorp/waypoint-client/gen/client/waypoint"
	wcs "github.com/hashicorp/waypoint-client/gen/client/waypoint_control_service"
)

const (
	defaultBasePath = "/"
)

type Config struct {
	//HCPConfig contains config values to interact with HCP Waypoint API
	config.HCPWaypointConfig

	//WaypointConfig contains config values to interact with Self Managed Waypoint API.
	config.WaypointConfig

	//BasePath is the path to the api server.
	BasePath string

	//APIState is the client state. SelfManaged for Self Managed Waypoint
	//and HCP for HCP Waypoint.
	APIState apiState
}

type apiState int

const (
	HCP apiState = iota
	SelfManaged
	Undetermined
)

// New creates a new Waypoint client given the config.
func New(config Config) (smwaypoint.ClientService, error) {

	apiState, err := config.decideState()
	if err != nil {
		return nil, fmt.Errorf("Cannot set APIState with current config.")
	}
	config.APIState = apiState

	tlsTransport := cleanhttp.DefaultPooledTransport()
	tlsTransport.TLSClientConfig = &tls.Config{}

	var transport http.RoundTripper = &oauth2.Transport{
		Base:   tlsTransport,
		Source: &config,
	}
	if config.BasePath == "" {
		config.BasePath = "/"
	}

	var bearerAuth runtime2.ClientAuthInfoWriter;
	if config.APIState == HCP {
		//create client to get namespace only. this client is only used once on initialization to grab the namespace.
		runtime := httptransport.New(config.APIAddress(), config.BasePath, []string{"https"})

		token, err := config.HCPWaypointConfig.Token()
		if err != nil {
			return nil, fmt.Errorf("Cannot retrieve token from current HCP Credentials.")
		}

		bearerAuth := httptransport.BearerToken(token.AccessToken)
		runtime.DefaultAuthentication = bearerAuth

		namespaceClient := wcs.New(runtime, strfmt.Default)
		//Do GetNamespace Call. Put it in the context.
		namespaceId, err := namespace.GetNamespace(namespaceClient, config.HCPOrgId, config.HCPProjectId)
		if err != nil {
			return nil, err
		}
		if namespaceId == "" || namespaceId == "0" {
			return nil, fmt.Errorf("No namespace found associated with orgId and projectId provided.")
		}

		basePath := config.BasePath + "namespace/" + namespaceId

		config.BasePath = basePath
	}

	runtime := httptransport.New(config.APIAddress(), config.BasePath, []string{"https"})
	runtime.Transport = transport
	runtime.DefaultAuthentication = bearerAuth

	apiClient := smwaypoint.New(runtime, strfmt.Default)

	return apiClient, nil
}

// namespaceWithContext generates a new context that contains the namespace identifiers
// currently in use.
func namespaceWithContext(ctx context.Context, externalId string, organizationId string, hcpProjectId string) context.Context {
	ctx = namespace.InContext(ctx, externalId, organizationId, hcpProjectId)

	// Also decorate the logger on the context with the namespace id
	return hclog.WithContext(ctx, hclog.FromContext(ctx), "namespace_id", externalId)
}

func (c *Config) decideState() (apiState, error) {
	err := c.HCPWaypointConfig.Validate()
	if err == nil {
		return HCP, nil
	}
	err = c.WaypointConfig.Validate()
	if err == nil {
		return SelfManaged, nil
	}
	return Undetermined, err
}

func (c *Config) APIAddress() string {
	if c.APIState == SelfManaged {
		return c.ServerUrl
	} else {
		return "api.cloud.hashicorp.com"
	}
}
