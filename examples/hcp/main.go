package main

import (
	"context"
	"fmt"
	"log"

	apiclient "github.com/hashicorp/waypoint-client/client"
	"github.com/hashicorp/waypoint/pkg/client/gen/client/waypoint"
)

//Example Usage of Waypoint Client Library with HCP Waypoint.

func main() {

	clientConfig, err := apiclient.NewWithHCP("hcp_org_id",
		"hcp_project_id",
		"client_id",
		"client_secret", context.Background())
	if err != nil {
		panic(err)
	}

	client, err := apiclient.New(*clientConfig)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.WaypointListProjects(waypoint.NewWaypointListProjectsParams())
	if err != nil {
		log.Fatal("An error occurred while fetching projects", err)
	}

	projects := resp.GetPayload().Projects

	for _, proj := range projects {
		fmt.Println(proj.Project)
	}
}
