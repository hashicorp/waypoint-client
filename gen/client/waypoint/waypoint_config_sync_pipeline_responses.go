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

// WaypointConfigSyncPipelineReader is a Reader for the WaypointConfigSyncPipeline structure.
type WaypointConfigSyncPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointConfigSyncPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointConfigSyncPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointConfigSyncPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointConfigSyncPipelineOK creates a WaypointConfigSyncPipelineOK with default headers values
func NewWaypointConfigSyncPipelineOK() *WaypointConfigSyncPipelineOK {
	return &WaypointConfigSyncPipelineOK{}
}

/*WaypointConfigSyncPipelineOK handles this case with default header values.

A successful response.
*/
type WaypointConfigSyncPipelineOK struct {
	Payload *models.HashicorpWaypointConfigSyncPipelineResponse
}

func (o *WaypointConfigSyncPipelineOK) Error() string {
	return fmt.Sprintf("[POST /project/{project.project}/config-sync-pipeline][%d] waypointConfigSyncPipelineOK  %+v", 200, o.Payload)
}

func (o *WaypointConfigSyncPipelineOK) GetPayload() *models.HashicorpWaypointConfigSyncPipelineResponse {
	return o.Payload
}

func (o *WaypointConfigSyncPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointConfigSyncPipelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointConfigSyncPipelineDefault creates a WaypointConfigSyncPipelineDefault with default headers values
func NewWaypointConfigSyncPipelineDefault(code int) *WaypointConfigSyncPipelineDefault {
	return &WaypointConfigSyncPipelineDefault{
		_statusCode: code,
	}
}

/*WaypointConfigSyncPipelineDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointConfigSyncPipelineDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint config sync pipeline default response
func (o *WaypointConfigSyncPipelineDefault) Code() int {
	return o._statusCode
}

func (o *WaypointConfigSyncPipelineDefault) Error() string {
	return fmt.Sprintf("[POST /project/{project.project}/config-sync-pipeline][%d] Waypoint_ConfigSyncPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointConfigSyncPipelineDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointConfigSyncPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}