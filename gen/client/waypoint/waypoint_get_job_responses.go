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

// WaypointGetJobReader is a Reader for the WaypointGetJob structure.
type WaypointGetJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetJobOK creates a WaypointGetJobOK with default headers values
func NewWaypointGetJobOK() *WaypointGetJobOK {
	return &WaypointGetJobOK{}
}

/*WaypointGetJobOK handles this case with default header values.

A successful response.
*/
type WaypointGetJobOK struct {
	Payload *models.HashicorpWaypointJob
}

func (o *WaypointGetJobOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/{job_id}][%d] waypointGetJobOK  %+v", 200, o.Payload)
}

func (o *WaypointGetJobOK) GetPayload() *models.HashicorpWaypointJob {
	return o.Payload
}

func (o *WaypointGetJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetJobDefault creates a WaypointGetJobDefault with default headers values
func NewWaypointGetJobDefault(code int) *WaypointGetJobDefault {
	return &WaypointGetJobDefault{
		_statusCode: code,
	}
}

/*WaypointGetJobDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetJobDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get job default response
func (o *WaypointGetJobDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetJobDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/jobs/{job_id}][%d] Waypoint_GetJob default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetJobDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
