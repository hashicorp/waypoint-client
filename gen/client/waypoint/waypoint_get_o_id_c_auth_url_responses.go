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

// WaypointGetOIDCAuthURLReader is a Reader for the WaypointGetOIDCAuthURL structure.
type WaypointGetOIDCAuthURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetOIDCAuthURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetOIDCAuthURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetOIDCAuthURLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetOIDCAuthURLOK creates a WaypointGetOIDCAuthURLOK with default headers values
func NewWaypointGetOIDCAuthURLOK() *WaypointGetOIDCAuthURLOK {
	return &WaypointGetOIDCAuthURLOK{}
}

/*WaypointGetOIDCAuthURLOK handles this case with default header values.

A successful response.
*/
type WaypointGetOIDCAuthURLOK struct {
	Payload *models.HashicorpWaypointGetOIDCAuthURLResponse
}

func (o *WaypointGetOIDCAuthURLOK) Error() string {
	return fmt.Sprintf("[POST /oidc/{auth_method.name}/url][%d] waypointGetOIdCAuthUrlOK  %+v", 200, o.Payload)
}

func (o *WaypointGetOIDCAuthURLOK) GetPayload() *models.HashicorpWaypointGetOIDCAuthURLResponse {
	return o.Payload
}

func (o *WaypointGetOIDCAuthURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetOIDCAuthURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetOIDCAuthURLDefault creates a WaypointGetOIDCAuthURLDefault with default headers values
func NewWaypointGetOIDCAuthURLDefault(code int) *WaypointGetOIDCAuthURLDefault {
	return &WaypointGetOIDCAuthURLDefault{
		_statusCode: code,
	}
}

/*WaypointGetOIDCAuthURLDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetOIDCAuthURLDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get o ID c auth URL default response
func (o *WaypointGetOIDCAuthURLDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetOIDCAuthURLDefault) Error() string {
	return fmt.Sprintf("[POST /oidc/{auth_method.name}/url][%d] Waypoint_GetOIDCAuthURL default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetOIDCAuthURLDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetOIDCAuthURLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}