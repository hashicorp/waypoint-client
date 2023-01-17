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

// WaypointListWorkspacesReader is a Reader for the WaypointListWorkspaces structure.
type WaypointListWorkspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListWorkspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListWorkspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListWorkspacesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListWorkspacesOK creates a WaypointListWorkspacesOK with default headers values
func NewWaypointListWorkspacesOK() *WaypointListWorkspacesOK {
	return &WaypointListWorkspacesOK{}
}

/*WaypointListWorkspacesOK handles this case with default header values.

A successful response.
*/
type WaypointListWorkspacesOK struct {
	Payload *models.HashicorpWaypointListWorkspacesResponse
}

func (o *WaypointListWorkspacesOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/workspaces][%d] waypointListWorkspacesOK  %+v", 200, o.Payload)
}

func (o *WaypointListWorkspacesOK) GetPayload() *models.HashicorpWaypointListWorkspacesResponse {
	return o.Payload
}

func (o *WaypointListWorkspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListWorkspacesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListWorkspacesDefault creates a WaypointListWorkspacesDefault with default headers values
func NewWaypointListWorkspacesDefault(code int) *WaypointListWorkspacesDefault {
	return &WaypointListWorkspacesDefault{
		_statusCode: code,
	}
}

/*WaypointListWorkspacesDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointListWorkspacesDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list workspaces default response
func (o *WaypointListWorkspacesDefault) Code() int {
	return o._statusCode
}

func (o *WaypointListWorkspacesDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/workspaces][%d] Waypoint_ListWorkspaces default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListWorkspacesDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListWorkspacesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
