# Waypoint Client

This API Client provides a Go Library for using Waypoint's Public API

# Installation
Fetch and install the package:

```bash
go get github.com/hashicorp/waypoint-client
```

## Authentication

This Waypoint Client can authenticate against Self-Managed Waypoint instances or HCP Waypoint instance. 

### HCP Waypoint

Client credentials are required for authentication against HCP Waypoint. 
In order to use this API Client with HCP Waypoint, you will need to provide the API Client 
with the following config values for authentication. Found in the HCP Portal https://portal.cloud.hashicorp.com/


#### HCPOrgId 
the organization id.

#### HCPProjectId
the hcp project id. 

#### ClientId
The application id. 

#### ClientSecret
The application secret. 



### Self Managed Waypoint
To authenticate against a Self Managed Instance of Waypoint, you will need to provide the API Client
with the following config values:


#### ServerUrl
The URL of the Waypoint Server.

#### WaypointUserToken
The user token used for auth in self-managed Waypoint.

#### InsecureSkipVerify
Waypoint by default runs with a self-signed certificate.

## Usage

1. Add the API Client Library to your go.mod

    ```go
        require (
            github.com/hashicorp/waypoint-client {latest release}
    ```

2. Example code for initializing a HCP Waypoint API Client

```go
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
```

3. Example code for initializing a Self Managed Waypoint API Client

```go
    clientConfig, err := client.NewWithSelfManaged("waypoint-user-token", "localhost:9702", true)
	if err != nil {
		panic(err)
	}
   
	client, err := client.New(*clientConfig)
	if err != nil {
         log.Fatal("Unable to setup client", err)
	}
```


See `examples/hcp/main.go` for a complete example for HCP and `examples/self-managed/main.go` for a self-managed Waypoint server.

