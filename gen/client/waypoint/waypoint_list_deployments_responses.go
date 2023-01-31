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

// WaypointListDeploymentsReader is a Reader for the WaypointListDeployments structure.
type WaypointListDeploymentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListDeploymentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListDeploymentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListDeploymentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListDeploymentsOK creates a WaypointListDeploymentsOK with default headers values
func NewWaypointListDeploymentsOK() *WaypointListDeploymentsOK {
	return &WaypointListDeploymentsOK{}
}

/*WaypointListDeploymentsOK handles this case with default header values.

A successful response.
*/
type WaypointListDeploymentsOK struct {
	Payload *models.HashicorpWaypointListDeploymentsResponse
}

func (o *WaypointListDeploymentsOK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/deployments][%d] waypointListDeploymentsOK  %+v", 200, o.Payload)
}

func (o *WaypointListDeploymentsOK) GetPayload() *models.HashicorpWaypointListDeploymentsResponse {
	return o.Payload
}

func (o *WaypointListDeploymentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListDeploymentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListDeploymentsDefault creates a WaypointListDeploymentsDefault with default headers values
func NewWaypointListDeploymentsDefault(code int) *WaypointListDeploymentsDefault {
	return &WaypointListDeploymentsDefault{
		_statusCode: code,
	}
}

/*WaypointListDeploymentsDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointListDeploymentsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list deployments default response
func (o *WaypointListDeploymentsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointListDeploymentsDefault) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/deployments][%d] Waypoint_ListDeployments default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListDeploymentsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListDeploymentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
