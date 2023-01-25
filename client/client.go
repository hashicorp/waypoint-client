package client

import (
	"context"
	"fmt"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	namespace "github.com/hashicorp/waypoint-client/context"
	"github.com/hashicorp/waypoint-client/gen/client"
)

type WaypointClientConfig struct {
	//HCPOrgId is the organization id. Only for HCP Waypoint.
	HCPOrgId string

	// HCPProjectId is the hcp project id. Only for HCP Waypoint.
	HCPProjectId string

	//clientCredentialsConfig is the service principal credentials that will be used for oauth.
	ClientCredentialsConfig clientcredentials.Config

	///WaypointUserToken is the user token used for auth in self-managed Waypoint. Only for Self-Managed Waypoint.
	WaypointUserToken string

	// ServerURL is the URL of the Waypoint Server. Only for Self-Managed Waypoint.
	ServerUrl string

	//BasePath is the path to the api server.
	BasePath string

	APIState apiState
}

type apiState int

const (
	HCP apiState = iota
	SelfManaged
)

// New creates a new Waypoint client given the config.
func New(config WaypointClientConfig) (*client.HashiCorpCloudPlatformWaypoint, error) {
	tlsTransport := cleanhttp.DefaultPooledTransport()

	var transport http.RoundTripper = &oauth2.Transport{
		Base: tlsTransport,
	}
	err := config.decideState()
	if err != nil {
		return nil, fmt.Errorf("Cannot set APIState with current config.")
	}

	ctx := context.Background()

	runtime := httptransport.New(config.APIAddress(), config.BasePath, []string{"https"})
	runtime.Transport = transport

	client := client.New(runtime, strfmt.Default)

	if config.APIState == HCP {
		//Do GetNamespace Call. Put it in the context.
		namespaceId, err := namespace.GetNamespace(client, config.HCPOrgId, config.HCPProjectId)
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

// TODO (teresa): Clean this up!
func (c *WaypointClientConfig) decideState() error {
	err := c.validateHCPConfig()
	//If there is an error on HCP env var validation, check if OSS config is valid.
	if err != nil {
		err = c.validateSelfManagedConfig()
		if err != nil {
			return err
		}
		c.BasePath = c.ServerUrl
	}
	c.BasePath = "api.cloud.hashicorp.com"
	return nil
}

func (c *WaypointClientConfig) validateHCPConfig() error {
	if (c.ClientCredentialsConfig.ClientID == "" && c.ClientCredentialsConfig.ClientSecret != "") ||
		(c.ClientCredentialsConfig.ClientID != "" && c.ClientCredentialsConfig.ClientSecret == "") {
		return fmt.Errorf("both client ID and secret must be provided")
	}

	if (c.HCPOrgId == "" && c.HCPProjectId != "") ||
		(c.HCPOrgId != "" && c.HCPProjectId == "") {
		return fmt.Errorf("when setting a user profile, both organization ID and project ID must be provided")
	}
	c.APIState = HCP

	return nil
}

func (c *WaypointClientConfig) validateSelfManagedConfig() error {
	if (c.WaypointUserToken == "" && c.ServerUrl != "") || (c.WaypointUserToken != "" && c.ServerUrl == "") {
		return fmt.Errorf("both user token and server url must be provided")
	}
	c.APIState = SelfManaged

	return nil
}

func (c *WaypointClientConfig) APIAddress() string {
	if c.APIState == SelfManaged {
		return c.ServerUrl
	} else {
		return "api.cloud.hashicorp.com"
	}
}
