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

// WaypointGetLatestBuildReader is a Reader for the WaypointGetLatestBuild structure.
type WaypointGetLatestBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetLatestBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetLatestBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetLatestBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetLatestBuildOK creates a WaypointGetLatestBuildOK with default headers values
func NewWaypointGetLatestBuildOK() *WaypointGetLatestBuildOK {
	return &WaypointGetLatestBuildOK{}
}

/*WaypointGetLatestBuildOK handles this case with default header values.

A successful response.
*/
type WaypointGetLatestBuildOK struct {
	Payload *models.HashicorpWaypointBuild
}

func (o *WaypointGetLatestBuildOK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/build/latest][%d] waypointGetLatestBuildOK  %+v", 200, o.Payload)
}

func (o *WaypointGetLatestBuildOK) GetPayload() *models.HashicorpWaypointBuild {
	return o.Payload
}

func (o *WaypointGetLatestBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointBuild)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetLatestBuildDefault creates a WaypointGetLatestBuildDefault with default headers values
func NewWaypointGetLatestBuildDefault(code int) *WaypointGetLatestBuildDefault {
	return &WaypointGetLatestBuildDefault{
		_statusCode: code,
	}
}

/*WaypointGetLatestBuildDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetLatestBuildDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get latest build default response
func (o *WaypointGetLatestBuildDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetLatestBuildDefault) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/build/latest][%d] Waypoint_GetLatestBuild default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetLatestBuildDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetLatestBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}