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

// WaypointRunPipelineReader is a Reader for the WaypointRunPipeline structure.
type WaypointRunPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointRunPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointRunPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointRunPipelineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointRunPipelineOK creates a WaypointRunPipelineOK with default headers values
func NewWaypointRunPipelineOK() *WaypointRunPipelineOK {
	return &WaypointRunPipelineOK{}
}

/*WaypointRunPipelineOK handles this case with default header values.

A successful response.
*/
type WaypointRunPipelineOK struct {
	Payload *models.HashicorpWaypointRunPipelineResponse
}

func (o *WaypointRunPipelineOK) Error() string {
	return fmt.Sprintf("[POST /project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/run][%d] waypointRunPipelineOK  %+v", 200, o.Payload)
}

func (o *WaypointRunPipelineOK) GetPayload() *models.HashicorpWaypointRunPipelineResponse {
	return o.Payload
}

func (o *WaypointRunPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointRunPipelineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointRunPipelineDefault creates a WaypointRunPipelineDefault with default headers values
func NewWaypointRunPipelineDefault(code int) *WaypointRunPipelineDefault {
	return &WaypointRunPipelineDefault{
		_statusCode: code,
	}
}

/*WaypointRunPipelineDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointRunPipelineDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint run pipeline default response
func (o *WaypointRunPipelineDefault) Code() int {
	return o._statusCode
}

func (o *WaypointRunPipelineDefault) Error() string {
	return fmt.Sprintf("[POST /project/{pipeline.owner.project.project}/pipeline/{pipeline.owner.pipeline_name}/run][%d] Waypoint_RunPipeline default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointRunPipelineDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointRunPipelineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}