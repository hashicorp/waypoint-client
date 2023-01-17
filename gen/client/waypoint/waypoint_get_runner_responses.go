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

// WaypointGetRunnerReader is a Reader for the WaypointGetRunner structure.
type WaypointGetRunnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetRunnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetRunnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetRunnerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetRunnerOK creates a WaypointGetRunnerOK with default headers values
func NewWaypointGetRunnerOK() *WaypointGetRunnerOK {
	return &WaypointGetRunnerOK{}
}

/*WaypointGetRunnerOK handles this case with default header values.

A successful response.
*/
type WaypointGetRunnerOK struct {
	Payload *models.HashicorpWaypointRunner
}

func (o *WaypointGetRunnerOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runner/{runner_id}][%d] waypointGetRunnerOK  %+v", 200, o.Payload)
}

func (o *WaypointGetRunnerOK) GetPayload() *models.HashicorpWaypointRunner {
	return o.Payload
}

func (o *WaypointGetRunnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointRunner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetRunnerDefault creates a WaypointGetRunnerDefault with default headers values
func NewWaypointGetRunnerDefault(code int) *WaypointGetRunnerDefault {
	return &WaypointGetRunnerDefault{
		_statusCode: code,
	}
}

/*WaypointGetRunnerDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetRunnerDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get runner default response
func (o *WaypointGetRunnerDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetRunnerDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-02-03/namespace/{namespace_id}/runner/{runner_id}][%d] Waypoint_GetRunner default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetRunnerDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetRunnerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
