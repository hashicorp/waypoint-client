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

// WaypointListDeployments2Reader is a Reader for the WaypointListDeployments2 structure.
type WaypointListDeployments2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListDeployments2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListDeployments2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListDeployments2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListDeployments2OK creates a WaypointListDeployments2OK with default headers values
func NewWaypointListDeployments2OK() *WaypointListDeployments2OK {
	return &WaypointListDeployments2OK{}
}

/*WaypointListDeployments2OK handles this case with default header values.

A successful response.
*/
type WaypointListDeployments2OK struct {
	Payload *models.HashicorpWaypointListDeploymentsResponse
}

func (o *WaypointListDeployments2OK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/deployments][%d] waypointListDeployments2OK  %+v", 200, o.Payload)
}

func (o *WaypointListDeployments2OK) GetPayload() *models.HashicorpWaypointListDeploymentsResponse {
	return o.Payload
}

func (o *WaypointListDeployments2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListDeploymentsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListDeployments2Default creates a WaypointListDeployments2Default with default headers values
func NewWaypointListDeployments2Default(code int) *WaypointListDeployments2Default {
	return &WaypointListDeployments2Default{
		_statusCode: code,
	}
}

/*WaypointListDeployments2Default handles this case with default header values.

An unexpected error response.
*/
type WaypointListDeployments2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list deployments2 default response
func (o *WaypointListDeployments2Default) Code() int {
	return o._statusCode
}

func (o *WaypointListDeployments2Default) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/workspace/{workspace.workspace}/deployments][%d] Waypoint_ListDeployments2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListDeployments2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListDeployments2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
