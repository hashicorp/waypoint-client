// Copyright (c) HashiCorp, Inc.

package client

import (
	"fmt"
)

type WaypointConfig struct {

	///WaypointUserToken is the user token used for auth in self-managed Waypoint. Only for Self-Managed Waypoint.
	WaypointUserToken string

	// Waypoint by default runs with a self-signed certificate.
	InsecureSkipVerify bool
}

func NewWithSelfManaged(waypointUserToken, serverUrl string, skipVerify bool) (*Config, error) {
	waypointcfg := &WaypointConfig{
		WaypointUserToken:  waypointUserToken,
		InsecureSkipVerify: skipVerify,
	}

	if serverUrl == "" {
		return nil, fmt.Errorf("cannot create self managed waypoint config, no server url provided")
	}

	err := waypointcfg.validateSelfManaged()
	if err != nil {
		return nil, fmt.Errorf("cannot create self managed waypoint config: %w", err)
	}

	return &Config{
		WaypointConfig: waypointcfg,
		ServerUrl:      serverUrl,
	}, nil
}

func (c *WaypointConfig) validateSelfManaged() error {
	if c.WaypointUserToken == "" {
		return fmt.Errorf("user token must be provided")
	}

	return nil
}
