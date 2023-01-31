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

// WaypointForgetRunnerReader is a Reader for the WaypointForgetRunner structure.
type WaypointForgetRunnerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointForgetRunnerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointForgetRunnerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointForgetRunnerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointForgetRunnerOK creates a WaypointForgetRunnerOK with default headers values
func NewWaypointForgetRunnerOK() *WaypointForgetRunnerOK {
	return &WaypointForgetRunnerOK{}
}

/*WaypointForgetRunnerOK handles this case with default header values.

A successful response.
*/
type WaypointForgetRunnerOK struct {
	Payload interface{}
}

func (o *WaypointForgetRunnerOK) Error() string {
	return fmt.Sprintf("[POST /runner/{runner_id}/forget][%d] waypointForgetRunnerOK  %+v", 200, o.Payload)
}

func (o *WaypointForgetRunnerOK) GetPayload() interface{} {
	return o.Payload
}

func (o *WaypointForgetRunnerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointForgetRunnerDefault creates a WaypointForgetRunnerDefault with default headers values
func NewWaypointForgetRunnerDefault(code int) *WaypointForgetRunnerDefault {
	return &WaypointForgetRunnerDefault{
		_statusCode: code,
	}
}

/*WaypointForgetRunnerDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointForgetRunnerDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint forget runner default response
func (o *WaypointForgetRunnerDefault) Code() int {
	return o._statusCode
}

func (o *WaypointForgetRunnerDefault) Error() string {
	return fmt.Sprintf("[POST /runner/{runner_id}/forget][%d] Waypoint_ForgetRunner default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointForgetRunnerDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointForgetRunnerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}