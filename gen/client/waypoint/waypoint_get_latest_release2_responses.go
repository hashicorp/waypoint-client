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

// WaypointGetLatestRelease2Reader is a Reader for the WaypointGetLatestRelease2 structure.
type WaypointGetLatestRelease2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetLatestRelease2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetLatestRelease2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetLatestRelease2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetLatestRelease2OK creates a WaypointGetLatestRelease2OK with default headers values
func NewWaypointGetLatestRelease2OK() *WaypointGetLatestRelease2OK {
	return &WaypointGetLatestRelease2OK{}
}

/*WaypointGetLatestRelease2OK handles this case with default header values.

A successful response.
*/
type WaypointGetLatestRelease2OK struct {
	Payload *models.HashicorpWaypointRelease
}

func (o *WaypointGetLatestRelease2OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/release/latest][%d] waypointGetLatestRelease2OK  %+v", 200, o.Payload)
}

func (o *WaypointGetLatestRelease2OK) GetPayload() *models.HashicorpWaypointRelease {
	return o.Payload
}

func (o *WaypointGetLatestRelease2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointRelease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetLatestRelease2Default creates a WaypointGetLatestRelease2Default with default headers values
func NewWaypointGetLatestRelease2Default(code int) *WaypointGetLatestRelease2Default {
	return &WaypointGetLatestRelease2Default{
		_statusCode: code,
	}
}

/*WaypointGetLatestRelease2Default handles this case with default header values.

An unexpected error response.
*/
type WaypointGetLatestRelease2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get latest release2 default response
func (o *WaypointGetLatestRelease2Default) Code() int {
	return o._statusCode
}

func (o *WaypointGetLatestRelease2Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/release/latest][%d] Waypoint_GetLatestRelease2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetLatestRelease2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetLatestRelease2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
