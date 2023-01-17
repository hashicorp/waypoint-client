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

// WaypointListOnDemandRunnerConfigsReader is a Reader for the WaypointListOnDemandRunnerConfigs structure.
type WaypointListOnDemandRunnerConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListOnDemandRunnerConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListOnDemandRunnerConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListOnDemandRunnerConfigsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListOnDemandRunnerConfigsOK creates a WaypointListOnDemandRunnerConfigsOK with default headers values
func NewWaypointListOnDemandRunnerConfigsOK() *WaypointListOnDemandRunnerConfigsOK {
	return &WaypointListOnDemandRunnerConfigsOK{}
}

/*WaypointListOnDemandRunnerConfigsOK handles this case with default header values.

A successful response.
*/
type WaypointListOnDemandRunnerConfigsOK struct {
	Payload *models.HashicorpWaypointListOnDemandRunnerConfigsResponse
}

func (o *WaypointListOnDemandRunnerConfigsOK) Error() string {
	return fmt.Sprintf("[GET /on-demand-runners][%d] waypointListOnDemandRunnerConfigsOK  %+v", 200, o.Payload)
}

func (o *WaypointListOnDemandRunnerConfigsOK) GetPayload() *models.HashicorpWaypointListOnDemandRunnerConfigsResponse {
	return o.Payload
}

func (o *WaypointListOnDemandRunnerConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListOnDemandRunnerConfigsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListOnDemandRunnerConfigsDefault creates a WaypointListOnDemandRunnerConfigsDefault with default headers values
func NewWaypointListOnDemandRunnerConfigsDefault(code int) *WaypointListOnDemandRunnerConfigsDefault {
	return &WaypointListOnDemandRunnerConfigsDefault{
		_statusCode: code,
	}
}

/*WaypointListOnDemandRunnerConfigsDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointListOnDemandRunnerConfigsDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list on demand runner configs default response
func (o *WaypointListOnDemandRunnerConfigsDefault) Code() int {
	return o._statusCode
}

func (o *WaypointListOnDemandRunnerConfigsDefault) Error() string {
	return fmt.Sprintf("[GET /on-demand-runners][%d] Waypoint_ListOnDemandRunnerConfigs default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListOnDemandRunnerConfigsDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListOnDemandRunnerConfigsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
