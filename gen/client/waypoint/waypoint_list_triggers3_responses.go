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

// WaypointListTriggers3Reader is a Reader for the WaypointListTriggers3 structure.
type WaypointListTriggers3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListTriggers3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListTriggers3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListTriggers3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListTriggers3OK creates a WaypointListTriggers3OK with default headers values
func NewWaypointListTriggers3OK() *WaypointListTriggers3OK {
	return &WaypointListTriggers3OK{}
}

/*WaypointListTriggers3OK handles this case with default header values.

A successful response.
*/
type WaypointListTriggers3OK struct {
	Payload *models.HashicorpWaypointListTriggerResponse
}

func (o *WaypointListTriggers3OK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{application.project}/application/{application.application}/triggers][%d] waypointListTriggers3OK  %+v", 200, o.Payload)
}

func (o *WaypointListTriggers3OK) GetPayload() *models.HashicorpWaypointListTriggerResponse {
	return o.Payload
}

func (o *WaypointListTriggers3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListTriggerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListTriggers3Default creates a WaypointListTriggers3Default with default headers values
func NewWaypointListTriggers3Default(code int) *WaypointListTriggers3Default {
	return &WaypointListTriggers3Default{
		_statusCode: code,
	}
}

/*WaypointListTriggers3Default handles this case with default header values.

An unexpected error response.
*/
type WaypointListTriggers3Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list triggers3 default response
func (o *WaypointListTriggers3Default) Code() int {
	return o._statusCode
}

func (o *WaypointListTriggers3Default) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/project/{application.project}/application/{application.application}/triggers][%d] Waypoint_ListTriggers3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListTriggers3Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListTriggers3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
