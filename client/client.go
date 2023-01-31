package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"

	"github.com/hashicorp/waypoint-client/config"
	namespace "github.com/hashicorp/waypoint-client/context"
	"github.com/hashicorp/waypoint-client/gen/client"
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
func New(config Config) (*client.HashiCorpCloudPlatformWaypoint, error) {
	tlsTransport := cleanhttp.DefaultPooledTransport()
	tlsTransport.TLSClientConfig = &tls.Config{}

	var transport http.RoundTripper = &oauth2.Transport{
		Base:   tlsTransport,
		Source: &config,
	}
	apiState, err := config.decideState()
	if err != nil {
		return nil, fmt.Errorf("Cannot set APIState with current config.")
	}
	config.APIState = apiState

	ctx := context.Background()

	runtime := httptransport.New(config.APIAddress(), config.BasePath, []string{"https"})
	runtime.Transport = transport

	if config.APIState == HCP {
		//create client to get namespace only. this client is only used once on initialization to grab the namespace.
		namespaceClient := client.New(runtime, strfmt.Default)

		token, err := config.HCPWaypointConfig.Token()
		if err != nil {
			return nil, fmt.Errorf("Cannot retrieve token from current HCP Credentials.")
		}

		bearerAuth := httptransport.BearerToken(token.AccessToken)

		//Do GetNamespace Call. Put it in the context.
		namespaceId, err := namespace.GetNamespace(namespaceClient, config.HCPOrgId, config.HCPProjectId, bearerAuth)
		if err != nil {
			return nil, err
		}
		namespaceWithContext(ctx, namespaceId, config.HCPOrgId, config.HCPProjectId)

	}

	return client, nil
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
