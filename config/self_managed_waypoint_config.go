package config

import "fmt"

type WaypointConfig struct {

	///WaypointUserToken is the user token used for auth in self-managed Waypoint. Only for Self-Managed Waypoint.
	WaypointUserToken string

	// ServerURL is the URL of the Waypoint Server. Only for Self-Managed Waypoint.
	ServerUrl string
}

func (c *WaypointConfig) Validate() error {
	if c == nil {
		return fmt.Errorf("Self managed config undefined")
	}

	if c.WaypointUserToken == "" || c.ServerUrl == "" {
		return fmt.Errorf("both user token and server url must be provided")
	}

	return nil
}
