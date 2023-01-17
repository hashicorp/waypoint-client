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

// WaypointListRunnersReader is a Reader for the WaypointListRunners structure.
type WaypointListRunnersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListRunnersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListRunnersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListRunnersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListRunnersOK creates a WaypointListRunnersOK with default headers values
func NewWaypointListRunnersOK() *WaypointListRunnersOK {
	return &WaypointListRunnersOK{}
}

/*WaypointListRunnersOK handles this case with default header values.

A successful response.
*/
type WaypointListRunnersOK struct {
	Payload *models.HashicorpWaypointListRunnersResponse
}

func (o *WaypointListRunnersOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runners][%d] waypointListRunnersOK  %+v", 200, o.Payload)
}

func (o *WaypointListRunnersOK) GetPayload() *models.HashicorpWaypointListRunnersResponse {
	return o.Payload
}

func (o *WaypointListRunnersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListRunnersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListRunnersDefault creates a WaypointListRunnersDefault with default headers values
func NewWaypointListRunnersDefault(code int) *WaypointListRunnersDefault {
	return &WaypointListRunnersDefault{
		_statusCode: code,
	}
}

/*WaypointListRunnersDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointListRunnersDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list runners default response
func (o *WaypointListRunnersDefault) Code() int {
	return o._statusCode
}

func (o *WaypointListRunnersDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runners][%d] Waypoint_ListRunners default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListRunnersDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListRunnersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
