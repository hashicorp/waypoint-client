// Copyright (c) HashiCorp, Inc.

package client

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"

	wcs "github.com/hashicorp/waypoint-client/gen/client/waypoint_control_service"
	smwaypoint "github.com/hashicorp/waypoint/pkg/client/gen/client/waypoint"
)

const (
	defaultBasePath = "/v1"
	hcpPathPrefix   = "waypoint/2022-02-03/namespace/"
)

type Config struct {
	//HCPConfig contains config values to interact with HCP Waypoint API
	//This config should be nil if interacting with Self Managed Waypoint API
	*HCPWaypointConfig

	//WaypointConfig contains config values to interact with Self Managed Waypoint API.
	//This config should be set to nil if interacting with HCP Waypoint API.
	*WaypointConfig

	//BasePath is the path to the api server.
	BasePath string

	//ServerURL is the URL of the Waypoint Server. Only for Self-Managed Waypoint.
	ServerUrl string

	//Schemes is the protocol the client will be set to use. Defaults to https
	Schemes []string
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

	var authWriter runtime.ClientAuthInfoWriter
	var tlsOpts httptransport.TLSClientOptions
	var httpRuntime *client.Runtime

	if len(config.Schemes) == 0 {
		config.Schemes = []string{"https"}
	}

	if config.BasePath == "" {
		config.BasePath = defaultBasePath
	}

	if config.isHCP() {
		//In order for the client to reach HCP Waypoint, we initialize a one-use only namespace client to make a request to
		//GetNamespace based on the provided orgId and hcpProjectId. It utilizes the waypoint control service api to make this request.
		//This GetNamespace call only happens on API client initialization for HCP Waypoint only.

		token, err := config.HCPWaypointConfig.Token()
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve token from current HCPWaypointConfig values: %w", err)
		}

		//create client to get namespace only. this client is only used once on initialization to grab the namespace.
		httpRuntime = httptransport.New(config.ServerUrl, config.BasePath, config.Schemes)

		authWriter = httptransport.BearerToken(token.AccessToken)
		httpRuntime.DefaultAuthentication = authWriter

		namespaceClient := wcs.New(httpRuntime, strfmt.Default)
		//Do GetNamespace Call. Put it in the context.

		if config.HCPOrgId == "" || config.HCPProjectId == "" {
			return nil, fmt.Errorf("orgId and HCPProjectId must be provided in HCPWaypointConfig. " +
				"Please login to https://portal.cloud.hashicorp.com/sign-in to retrieve these values")
		}

		resp, err := namespaceClient.WaypointControlServiceGetNamespace(
			wcs.NewWaypointControlServiceGetNamespaceParams().
				WithLocationOrganizationID(config.HCPOrgId).
				WithLocationProjectID(config.HCPProjectId),
			nil,
		)
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve namespace for orgId and project id: %w", err)
		}

		namespaceId := resp.GetPayload().Namespace.ID

		if namespaceId == "" || namespaceId == "0" {
			return nil, fmt.Errorf("no namespace found associated with orgId and projectId provided")
		}

		// Add namespace Id to the base path
		basePath := config.BasePath + hcpPathPrefix + namespaceId
		config.BasePath = basePath

	} else if config.isSelfManaged() {
		if config.ServerUrl == "" {
			return nil, fmt.Errorf("server url must be provided")
		}
		authWriter = httptransport.APIKeyAuth("authorization", "header", config.WaypointConfig.WaypointUserToken)

		tlsOpts = httptransport.TLSClientOptions{
			InsecureSkipVerify: config.InsecureSkipVerify,
		}
	}

	httpClient, err := httptransport.TLSClient(tlsOpts)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize http.Client: %w", err)
	}

	//Initialize API Client.
	httpRuntime = httptransport.NewWithClient(config.ServerUrl, config.BasePath, config.Schemes, httpClient)
	httpRuntime.Transport = transport
	httpRuntime.DefaultAuthentication = authWriter

	apiClient := smwaypoint.New(httpRuntime, strfmt.Default)

	return apiClient, nil
}
