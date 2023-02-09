package main

import (
	"fmt"
	"log"

	apiclient "github.com/hashicorp/waypoint-client/client"
	config "github.com/hashicorp/waypoint-client/config"
	"github.com/hashicorp/waypoint/pkg/client/gen/client/waypoint"
)

//Example Usage of Waypoint Client Library with HCP Waypoint.

func main() {

	hcp, err := config.New("some_hcp_org_id",
		"some_hcp_project_id",
		"some_service_principal_client_id",
		"some_client_secret")
	if err != nil {
		panic(err)
	}

	clientConfig := apiclient.Config{
		HCPWaypointConfig: *hcp,
		BasePath:          "/",
	}

	client, err := apiclient.New(clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.WaypointListProjects(waypoint.NewWaypointListProjectsParams())

	fmt.Printf("%v#", resp.GetPayload())

}
