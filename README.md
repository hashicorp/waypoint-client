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
with the following credentials found in the HCP Portal:

```bash
CLIENT_ID="..."
CLIENT_SECRET="..."
HCP_ORG_ID="..."
HCP_PROJECT_ID="..."
```

### Self Managed Waypoint
To authenticate against a Self Managed Instance of Waypoint, you will need to provide the API Client
with the following credentials:

```bash
SERVER_URL="..."
WAYPOINT_USER_TOKEN="..."
```

## Usage

1. Add the API Client Library to your go.mod

    ```go
    require (
        github.com/hashicorp/waypoint-client {latest release}
    ```

See `examples/hcp/main.go` for a complete example for HCP and `examples/self-managed/main.go` for a self-managed Waypoint server.

