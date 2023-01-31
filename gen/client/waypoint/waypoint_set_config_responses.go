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

// WaypointSetConfigReader is a Reader for the WaypointSetConfig structure.
type WaypointSetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointSetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointSetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointSetConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointSetConfigOK creates a WaypointSetConfigOK with default headers values
func NewWaypointSetConfigOK() *WaypointSetConfigOK {
	return &WaypointSetConfigOK{}
}

/*WaypointSetConfigOK handles this case with default header values.

A successful response.
*/
type WaypointSetConfigOK struct {
	Payload models.HashicorpWaypointConfigSetResponse
}

func (o *WaypointSetConfigOK) Error() string {
	return fmt.Sprintf("[PUT /project/config][%d] waypointSetConfigOK  %+v", 200, o.Payload)
}

func (o *WaypointSetConfigOK) GetPayload() models.HashicorpWaypointConfigSetResponse {
	return o.Payload
}

func (o *WaypointSetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointSetConfigDefault creates a WaypointSetConfigDefault with default headers values
func NewWaypointSetConfigDefault(code int) *WaypointSetConfigDefault {
	return &WaypointSetConfigDefault{
		_statusCode: code,
	}
}

/*WaypointSetConfigDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointSetConfigDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint set config default response
func (o *WaypointSetConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointSetConfigDefault) Error() string {
	return fmt.Sprintf("[PUT /project/config][%d] Waypoint_SetConfig default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointSetConfigDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointSetConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}