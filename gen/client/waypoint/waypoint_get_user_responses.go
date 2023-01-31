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

// WaypointGetUserReader is a Reader for the WaypointGetUser structure.
type WaypointGetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetUserOK creates a WaypointGetUserOK with default headers values
func NewWaypointGetUserOK() *WaypointGetUserOK {
	return &WaypointGetUserOK{}
}

/*WaypointGetUserOK handles this case with default header values.

A successful response.
*/
type WaypointGetUserOK struct {
	Payload *models.HashicorpWaypointGetUserResponse
}

func (o *WaypointGetUserOK) Error() string {
	return fmt.Sprintf("[GET /user/by-id/{user.id.id}][%d] waypointGetUserOK  %+v", 200, o.Payload)
}

func (o *WaypointGetUserOK) GetPayload() *models.HashicorpWaypointGetUserResponse {
	return o.Payload
}

func (o *WaypointGetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetUserDefault creates a WaypointGetUserDefault with default headers values
func NewWaypointGetUserDefault(code int) *WaypointGetUserDefault {
	return &WaypointGetUserDefault{
		_statusCode: code,
	}
}

/*WaypointGetUserDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetUserDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get user default response
func (o *WaypointGetUserDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetUserDefault) Error() string {
	return fmt.Sprintf("[GET /user/by-id/{user.id.id}][%d] Waypoint_GetUser default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetUserDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
