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

// WaypointExpediteStatusReport4Reader is a Reader for the WaypointExpediteStatusReport4 structure.
type WaypointExpediteStatusReport4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointExpediteStatusReport4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointExpediteStatusReport4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointExpediteStatusReport4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointExpediteStatusReport4OK creates a WaypointExpediteStatusReport4OK with default headers values
func NewWaypointExpediteStatusReport4OK() *WaypointExpediteStatusReport4OK {
	return &WaypointExpediteStatusReport4OK{}
}

/*WaypointExpediteStatusReport4OK handles this case with default header values.

A successful response.
*/
type WaypointExpediteStatusReport4OK struct {
	Payload *models.HashicorpWaypointExpediteStatusReportResponse
}

func (o *WaypointExpediteStatusReport4OK) Error() string {
	return fmt.Sprintf("[PUT /project/{release.sequence.application.project}/application/{release.sequence.application.application}/release/{release.sequence.number}/status-report][%d] waypointExpediteStatusReport4OK  %+v", 200, o.Payload)
}

func (o *WaypointExpediteStatusReport4OK) GetPayload() *models.HashicorpWaypointExpediteStatusReportResponse {
	return o.Payload
}

func (o *WaypointExpediteStatusReport4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointExpediteStatusReportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointExpediteStatusReport4Default creates a WaypointExpediteStatusReport4Default with default headers values
func NewWaypointExpediteStatusReport4Default(code int) *WaypointExpediteStatusReport4Default {
	return &WaypointExpediteStatusReport4Default{
		_statusCode: code,
	}
}

/*WaypointExpediteStatusReport4Default handles this case with default header values.

An unexpected error response.
*/
type WaypointExpediteStatusReport4Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint expedite status report4 default response
func (o *WaypointExpediteStatusReport4Default) Code() int {
	return o._statusCode
}

func (o *WaypointExpediteStatusReport4Default) Error() string {
	return fmt.Sprintf("[PUT /project/{release.sequence.application.project}/application/{release.sequence.application.application}/release/{release.sequence.number}/status-report][%d] Waypoint_ExpediteStatusReport4 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointExpediteStatusReport4Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointExpediteStatusReport4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}