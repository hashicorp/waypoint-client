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

// WaypointExpediteStatusReport3Reader is a Reader for the WaypointExpediteStatusReport3 structure.
type WaypointExpediteStatusReport3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointExpediteStatusReport3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointExpediteStatusReport3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointExpediteStatusReport3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointExpediteStatusReport3OK creates a WaypointExpediteStatusReport3OK with default headers values
func NewWaypointExpediteStatusReport3OK() *WaypointExpediteStatusReport3OK {
	return &WaypointExpediteStatusReport3OK{}
}

/*WaypointExpediteStatusReport3OK handles this case with default header values.

A successful response.
*/
type WaypointExpediteStatusReport3OK struct {
	Payload *models.HashicorpWaypointExpediteStatusReportResponse
}

func (o *WaypointExpediteStatusReport3OK) Error() string {
	return fmt.Sprintf("[PUT /project/{deployment.sequence.application.project}/application/{deployment.sequence.application.application}/deployment/{deployment.sequence.number}/status-report][%d] waypointExpediteStatusReport3OK  %+v", 200, o.Payload)
}

func (o *WaypointExpediteStatusReport3OK) GetPayload() *models.HashicorpWaypointExpediteStatusReportResponse {
	return o.Payload
}

func (o *WaypointExpediteStatusReport3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointExpediteStatusReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointExpediteStatusReport3Default creates a WaypointExpediteStatusReport3Default with default headers values
func NewWaypointExpediteStatusReport3Default(code int) *WaypointExpediteStatusReport3Default {
	return &WaypointExpediteStatusReport3Default{
		_statusCode: code,
	}
}

/*WaypointExpediteStatusReport3Default handles this case with default header values.

An unexpected error response.
*/
type WaypointExpediteStatusReport3Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint expedite status report3 default response
func (o *WaypointExpediteStatusReport3Default) Code() int {
	return o._statusCode
}

func (o *WaypointExpediteStatusReport3Default) Error() string {
	return fmt.Sprintf("[PUT /project/{deployment.sequence.application.project}/application/{deployment.sequence.application.application}/deployment/{deployment.sequence.number}/status-report][%d] Waypoint_ExpediteStatusReport3 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointExpediteStatusReport3Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointExpediteStatusReport3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
