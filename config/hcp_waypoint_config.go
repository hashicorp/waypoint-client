package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	// defaultAuthURL is the URL of the production auth endpoint.
	defaultAuthURL = "https://auth.idp.hashicorp.com"

	// The audience is the API identifier configured in the auth provider and
	// must be provided when requesting an access token for the API. The value
	// is the same regardless of environment.
	aud = "https://api.hashicorp.cloud"

	// tokenPath is the path used to retrieve the access token.
	tokenPath string = "/oauth2/token"
)

type HCPWaypointConfig struct {
	//HCPOrgId is the organization id. Only for HCP Waypoint.
	HCPOrgId string

	// HCPProjectId is the hcp project id. Only for HCP Waypoint.
	HCPProjectId string

	//clientCredentialsConfig is the service principal credentials that will be used for oauth.
	ClientCredentialsConfig clientcredentials.Config

	//BasePath is the path to the api server.
	BasePath string

	// tokenSource is a self-refreshing token source that returns an access
	// token that can be used to authenticate against the HCP APIs.
	tokenSource oauth2.TokenSource
}

// New creates a new HCP Waypoint Configuration with bearer token auth given HCP Waypoint parameters.
func New(orgId string, projectId string, clientId string, clientSecret string, ctx context.Context) (*HCPWaypointConfig, error) {
	tlsTransport := cleanhttp.DefaultTransport()
	tlsTransport.TLSClientConfig = &tls.Config{}

	config := &HCPWaypointConfig{
		HCPOrgId:     orgId,
		HCPProjectId: projectId,
		ClientCredentialsConfig: clientcredentials.Config{
			ClientID:       clientId,
			ClientSecret:   clientSecret,
			TokenURL:       defaultAuthURL,
			EndpointParams: url.Values{"audience": {aud}},
		},
	}

	tokenURL, err := url.Parse(defaultAuthURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse auth URL: %w", err)
	}

	tokenContext := context.WithValue(
		ctx,
		oauth2.HTTPClient,
		&http.Client{Transport: tlsTransport},
	)

	tokenURL.Path = tokenPath

	config.ClientCredentialsConfig.TokenURL = tokenURL.String()

	config.tokenSource = config.ClientCredentialsConfig.TokenSource(tokenContext)

	return config, nil
}

func (c *HCPWaypointConfig) Validate() error {
	if c.ClientCredentialsConfig.ClientID == "" || c.ClientCredentialsConfig.ClientSecret == "" ||
		c.HCPOrgId == "" || c.HCPProjectId == "" {
		return fmt.Errorf("one of the following fields is missing:" +
			"ClientID, " +
			"ClientSecret, " +
			"HCPOrgId, " +
			"HCPProjectId.")
	}
	return nil
}

func (c *HCPWaypointConfig) Token() (*oauth2.Token, error) {
	if c == nil {
		return nil, fmt.Errorf("Undefined")
	}
	return c.tokenSource.Token()
}
