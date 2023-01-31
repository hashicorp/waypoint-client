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

// WaypointListJobs4Reader is a Reader for the WaypointListJobs4 structure.
type WaypointListJobs4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListJobs4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListJobs4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListJobs4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListJobs4OK creates a WaypointListJobs4OK with default headers values
func NewWaypointListJobs4OK() *WaypointListJobs4OK {
	return &WaypointListJobs4OK{}
}

/*WaypointListJobs4OK handles this case with default header values.

A successful response.
*/
type WaypointListJobs4OK struct {
	Payload *models.HashicorpWaypointListJobsResponse
}

func (o *WaypointListJobs4OK) Error() string {
	return fmt.Sprintf("[GET /jobs/state/{jobState}][%d] waypointListJobs4OK  %+v", 200, o.Payload)
}

func (o *WaypointListJobs4OK) GetPayload() *models.HashicorpWaypointListJobsResponse {
	return o.Payload
}

func (o *WaypointListJobs4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListJobs4Default creates a WaypointListJobs4Default with default headers values
func NewWaypointListJobs4Default(code int) *WaypointListJobs4Default {
	return &WaypointListJobs4Default{
		_statusCode: code,
	}
}

/*WaypointListJobs4Default handles this case with default header values.

An unexpected error response.
*/
type WaypointListJobs4Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list jobs4 default response
func (o *WaypointListJobs4Default) Code() int {
	return o._statusCode
}

func (o *WaypointListJobs4Default) Error() string {
	return fmt.Sprintf("[GET /jobs/state/{jobState}][%d] Waypoint_ListJobs4 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListJobs4Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListJobs4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}