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

// WaypointUpsertAuthMethodReader is a Reader for the WaypointUpsertAuthMethod structure.
type WaypointUpsertAuthMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointUpsertAuthMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointUpsertAuthMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointUpsertAuthMethodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointUpsertAuthMethodOK creates a WaypointUpsertAuthMethodOK with default headers values
func NewWaypointUpsertAuthMethodOK() *WaypointUpsertAuthMethodOK {
	return &WaypointUpsertAuthMethodOK{}
}

/*WaypointUpsertAuthMethodOK handles this case with default header values.

A successful response.
*/
type WaypointUpsertAuthMethodOK struct {
	Payload *models.HashicorpWaypointUpsertAuthMethodResponse
}

func (o *WaypointUpsertAuthMethodOK) Error() string {
	return fmt.Sprintf("[POST /auth-method][%d] waypointUpsertAuthMethodOK  %+v", 200, o.Payload)
}

func (o *WaypointUpsertAuthMethodOK) GetPayload() *models.HashicorpWaypointUpsertAuthMethodResponse {
	return o.Payload
}

func (o *WaypointUpsertAuthMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointUpsertAuthMethodResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointUpsertAuthMethodDefault creates a WaypointUpsertAuthMethodDefault with default headers values
func NewWaypointUpsertAuthMethodDefault(code int) *WaypointUpsertAuthMethodDefault {
	return &WaypointUpsertAuthMethodDefault{
		_statusCode: code,
	}
}

/*WaypointUpsertAuthMethodDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointUpsertAuthMethodDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint upsert auth method default response
func (o *WaypointUpsertAuthMethodDefault) Code() int {
	return o._statusCode
}

func (o *WaypointUpsertAuthMethodDefault) Error() string {
	return fmt.Sprintf("[POST /auth-method][%d] Waypoint_UpsertAuthMethod default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointUpsertAuthMethodDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointUpsertAuthMethodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
