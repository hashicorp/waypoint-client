package context

import (
	"context"

	"github.com/go-openapi/runtime"

	"github.com/hashicorp/waypoint-client/gen/client"
	"github.com/hashicorp/waypoint-client/gen/client/waypoint_control_service"
)

type Key struct{}
type Value struct {
	ExternalId     string
	OrganizationId string
	HcpProjectId   string
}

func FromContext(ctx context.Context) (Value, bool) {
	v, ok := ctx.Value(Key{}).(Value)
	return v, ok
}

func InContext(ctx context.Context, externalId string, organizationId string, hcpProjectId string) context.Context {
	return context.WithValue(ctx, Key{}, Value{ExternalId: externalId, OrganizationId: organizationId, HcpProjectId: hcpProjectId})
}

func GetNamespace(c *client.HashiCorpCloudPlatformWaypoint, orgId string, hcpProjectId string, bearerAuth runtime.ClientAuthInfoWriter) (string, error) {
	resp, err := c.WaypointControlService.WaypointControlServiceGetNamespace(
		waypoint_control_service.NewWaypointControlServiceGetNamespaceParams().
			WithLocationOrganizationID(orgId).
			WithLocationProjectID(hcpProjectId),
		bearerAuth,
	)
	return resp.GetPayload().Namespace.ID, err

}
