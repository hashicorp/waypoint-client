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

// WaypointGetOnDemandRunnerConfigReader is a Reader for the WaypointGetOnDemandRunnerConfig structure.
type WaypointGetOnDemandRunnerConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetOnDemandRunnerConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetOnDemandRunnerConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetOnDemandRunnerConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetOnDemandRunnerConfigOK creates a WaypointGetOnDemandRunnerConfigOK with default headers values
func NewWaypointGetOnDemandRunnerConfigOK() *WaypointGetOnDemandRunnerConfigOK {
	return &WaypointGetOnDemandRunnerConfigOK{}
}

/*WaypointGetOnDemandRunnerConfigOK handles this case with default header values.

A successful response.
*/
type WaypointGetOnDemandRunnerConfigOK struct {
	Payload *models.HashicorpWaypointGetOnDemandRunnerConfigResponse
}

func (o *WaypointGetOnDemandRunnerConfigOK) Error() string {
	return fmt.Sprintf("[GET /on-demand-runner/by-id/{config.id}][%d] waypointGetOnDemandRunnerConfigOK  %+v", 200, o.Payload)
}

func (o *WaypointGetOnDemandRunnerConfigOK) GetPayload() *models.HashicorpWaypointGetOnDemandRunnerConfigResponse {
	return o.Payload
}

func (o *WaypointGetOnDemandRunnerConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointGetOnDemandRunnerConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetOnDemandRunnerConfigDefault creates a WaypointGetOnDemandRunnerConfigDefault with default headers values
func NewWaypointGetOnDemandRunnerConfigDefault(code int) *WaypointGetOnDemandRunnerConfigDefault {
	return &WaypointGetOnDemandRunnerConfigDefault{
		_statusCode: code,
	}
}

/*WaypointGetOnDemandRunnerConfigDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointGetOnDemandRunnerConfigDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get on demand runner config default response
func (o *WaypointGetOnDemandRunnerConfigDefault) Code() int {
	return o._statusCode
}

func (o *WaypointGetOnDemandRunnerConfigDefault) Error() string {
	return fmt.Sprintf("[GET /on-demand-runner/by-id/{config.id}][%d] Waypoint_GetOnDemandRunnerConfig default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetOnDemandRunnerConfigDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetOnDemandRunnerConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}