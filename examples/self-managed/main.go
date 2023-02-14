// Copyright (c) HashiCorp, Inc.

package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/waypoint-client/client"
	"github.com/hashicorp/waypoint/pkg/client/gen/client/waypoint"
)

//Example Usage of Waypoint Client Library with Self-Managed Waypoint.

func main() {

	clientConfig, err := client.NewWithSelfManaged("waypoint-user-token", "localhost:9702", true)
	if err != nil {
		panic(err)
	}

	client, err := client.New(*clientConfig)
	if err != nil {
		log.Fatal("Unable to setup client", err)
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
