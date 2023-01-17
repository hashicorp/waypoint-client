// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint-client/gen/models"
)

// WaypointUIListReleases4Reader is a Reader for the WaypointUIListReleases4 structure.
type WaypointUIListReleases4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUIListReleases4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUIListReleases4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUIListReleases4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUIListReleases4OK creates a WaypointUIListReleases4OK with default headers values
func NewWaypointUIListReleases4OK() *WaypointUIListReleases4OK {
	return &WaypointUIListReleases4OK{}
}

/*WaypointUIListReleases4OK handles this case with default header values.

A successful response.
*/
type WaypointUIListReleases4OK struct {
	Payload *models.HashicorpWaypointUIListReleasesResponse
}

func (o *WaypointUIListReleases4OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/ui/releases/state/{physical_state}][%d] waypointUiListReleases4OK  %+v", 200, o.Payload)
}

func (o *WaypointUIListReleases4OK) GetPayload() *models.HashicorpWaypointUIListReleasesResponse {
	return o.Payload
}

func (o *WaypointUIListReleases4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUIListReleasesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUIListReleases4Default creates a WaypointUIListReleases4Default with default headers values
func NewWaypointUIListReleases4Default(code int) *WaypointUIListReleases4Default {
	return &WaypointUIListReleases4Default{
		_statusCode: code,
	}
}

/*WaypointUIListReleases4Default handles this case with default header values.

An unexpected error response.
*/
type WaypointUIListReleases4Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint UI list releases4 default response
func (o *WaypointUIListReleases4Default) Code() int {
	return o._statusCode
}

func (o *WaypointUIListReleases4Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/ui/releases/state/{physical_state}][%d] Waypoint_UI_ListReleases4 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUIListReleases4Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUIListReleases4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
