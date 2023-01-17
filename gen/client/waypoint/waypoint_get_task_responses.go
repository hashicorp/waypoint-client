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

// WaypointGetTaskReader is a Reader for the WaypointGetTask structure.
type WaypointGetTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetTaskOK creates a WaypointGetTaskOK with default headers values
func NewWaypointGetTaskOK() *WaypointGetTaskOK {
	return &WaypointGetTaskOK{}
}

/*WaypointGetTaskOK handles this case with default header values.

A successful response.
*/
type WaypointGetTaskOK struct {
	Payload *models.HashicorpWaypointGetTaskResponse
}

func (o *WaypointGetTaskOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/task/{ref.id}][%d] waypointGetTaskOK  %+v", 200, o.Payload)
}

func (o *WaypointGetTaskOK) GetPayload() *models.HashicorpWaypointGetTaskResponse {
	return o.Payload
}

func (o *WaypointGetTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetTaskResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetTaskDefault creates a WaypointGetTaskDefault with default headers values
func NewWaypointGetTaskDefault(code int) *WaypointGetTaskDefault {
	return &WaypointGetTaskDefault{
		_statusCode: code,
	}
}

/*WaypointGetTaskDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetTaskDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get task default response
func (o *WaypointGetTaskDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetTaskDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/task/{ref.id}][%d] Waypoint_GetTask default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetTaskDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
