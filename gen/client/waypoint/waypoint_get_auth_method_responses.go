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

// WaypointGetAuthMethodReader is a Reader for the WaypointGetAuthMethod structure.
type WaypointGetAuthMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetAuthMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetAuthMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetAuthMethodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetAuthMethodOK creates a WaypointGetAuthMethodOK with default headers values
func NewWaypointGetAuthMethodOK() *WaypointGetAuthMethodOK {
	return &WaypointGetAuthMethodOK{}
}

/*WaypointGetAuthMethodOK handles this case with default header values.

A successful response.
*/
type WaypointGetAuthMethodOK struct {
	Payload *models.HashicorpWaypointGetAuthMethodResponse
}

func (o *WaypointGetAuthMethodOK) Error() string {
	return fmt.Sprintf("[GET /auth-method/{auth_method.name}][%d] waypointGetAuthMethodOK  %+v", 200, o.Payload)
}

func (o *WaypointGetAuthMethodOK) GetPayload() *models.HashicorpWaypointGetAuthMethodResponse {
	return o.Payload
}

func (o *WaypointGetAuthMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetAuthMethodResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetAuthMethodDefault creates a WaypointGetAuthMethodDefault with default headers values
func NewWaypointGetAuthMethodDefault(code int) *WaypointGetAuthMethodDefault {
	return &WaypointGetAuthMethodDefault{
		_statusCode: code,
	}
}

/*WaypointGetAuthMethodDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetAuthMethodDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get auth method default response
func (o *WaypointGetAuthMethodDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetAuthMethodDefault) Error() string {
	return fmt.Sprintf("[GET /auth-method/{auth_method.name}][%d] Waypoint_GetAuthMethod default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetAuthMethodDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetAuthMethodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}