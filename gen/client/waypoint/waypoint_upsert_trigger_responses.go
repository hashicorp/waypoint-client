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

// WaypointUpsertTriggerReader is a Reader for the WaypointUpsertTrigger structure.
type WaypointUpsertTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUpsertTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUpsertTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUpsertTriggerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUpsertTriggerOK creates a WaypointUpsertTriggerOK with default headers values
func NewWaypointUpsertTriggerOK() *WaypointUpsertTriggerOK {
	return &WaypointUpsertTriggerOK{}
}

/*WaypointUpsertTriggerOK handles this case with default header values.

A successful response.
*/
type WaypointUpsertTriggerOK struct {
	Payload *models.HashicorpWaypointUpsertTriggerResponse
}

func (o *WaypointUpsertTriggerOK) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/triggers][%d] waypointUpsertTriggerOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertTriggerOK) GetPayload() *models.HashicorpWaypointUpsertTriggerResponse {
	return o.Payload
}

func (o *WaypointUpsertTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUpsertTriggerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUpsertTriggerDefault creates a WaypointUpsertTriggerDefault with default headers values
func NewWaypointUpsertTriggerDefault(code int) *WaypointUpsertTriggerDefault {
	return &WaypointUpsertTriggerDefault{
		_statusCode: code,
	}
}

/*WaypointUpsertTriggerDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointUpsertTriggerDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint upsert trigger default response
func (o *WaypointUpsertTriggerDefault) Code() int {
	return o._statusCode
}

func (o *WaypointUpsertTriggerDefault) Error() string {
	return fmt.Sprintf("[POST /waypoint/2022-02-03/namespace/{namespace_id}/triggers][%d] Waypoint_UpsertTrigger default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertTriggerDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUpsertTriggerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
