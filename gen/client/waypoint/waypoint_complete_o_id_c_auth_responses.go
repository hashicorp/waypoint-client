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

// WaypointCompleteOIDCAuthReader is a Reader for the WaypointCompleteOIDCAuth structure.
type WaypointCompleteOIDCAuthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointCompleteOIDCAuthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointCompleteOIDCAuthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointCompleteOIDCAuthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointCompleteOIDCAuthOK creates a WaypointCompleteOIDCAuthOK with default headers values
func NewWaypointCompleteOIDCAuthOK() *WaypointCompleteOIDCAuthOK {
	return &WaypointCompleteOIDCAuthOK{}
}

/*WaypointCompleteOIDCAuthOK handles this case with default header values.

A successful response.
*/
type WaypointCompleteOIDCAuthOK struct {
	Payload *models.HashicorpWaypointCompleteOIDCAuthResponse
}

func (o *WaypointCompleteOIDCAuthOK) Error() string {
	return fmt.Sprintf("[POST /oidc/{auth_method.name}/complete][%d] waypointCompleteOIdCAuthOK  %+v", 200, o.Payload)
}

func (o *WaypointCompleteOIDCAuthOK) GetPayload() *models.HashicorpWaypointCompleteOIDCAuthResponse {
	return o.Payload
}

func (o *WaypointCompleteOIDCAuthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointCompleteOIDCAuthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointCompleteOIDCAuthDefault creates a WaypointCompleteOIDCAuthDefault with default headers values
func NewWaypointCompleteOIDCAuthDefault(code int) *WaypointCompleteOIDCAuthDefault {
	return &WaypointCompleteOIDCAuthDefault{
		_statusCode: code,
	}
}

/*WaypointCompleteOIDCAuthDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointCompleteOIDCAuthDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint complete o ID c auth default response
func (o *WaypointCompleteOIDCAuthDefault) Code() int {
	return o._statusCode
}

func (o *WaypointCompleteOIDCAuthDefault) Error() string {
	return fmt.Sprintf("[POST /oidc/{auth_method.name}/complete][%d] Waypoint_CompleteOIDCAuth default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointCompleteOIDCAuthDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointCompleteOIDCAuthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
