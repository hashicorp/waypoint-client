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

// WaypointListReleasesReader is a Reader for the WaypointListReleases structure.
type WaypointListReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointListReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointListReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointListReleasesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointListReleasesOK creates a WaypointListReleasesOK with default headers values
func NewWaypointListReleasesOK() *WaypointListReleasesOK {
	return &WaypointListReleasesOK{}
}

/*WaypointListReleasesOK handles this case with default header values.

A successful response.
*/
type WaypointListReleasesOK struct {
	Payload *models.HashicorpWaypointListReleasesResponse
}

func (o *WaypointListReleasesOK) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/releases][%d] waypointListReleasesOK  %+v", 200, o.Payload)
}

func (o *WaypointListReleasesOK) GetPayload() *models.HashicorpWaypointListReleasesResponse {
	return o.Payload
}

func (o *WaypointListReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointListReleasesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointListReleasesDefault creates a WaypointListReleasesDefault with default headers values
func NewWaypointListReleasesDefault(code int) *WaypointListReleasesDefault {
	return &WaypointListReleasesDefault{
		_statusCode: code,
	}
}

/*WaypointListReleasesDefault handles this case with default header values.

An unexpected error response.
*/
type WaypointListReleasesDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint list releases default response
func (o *WaypointListReleasesDefault) Code() int {
	return o._statusCode
}

func (o *WaypointListReleasesDefault) Error() string {
	return fmt.Sprintf("[GET /project/{application.project}/application/{application.application}/releases][%d] Waypoint_ListReleases default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointListReleasesDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointListReleasesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
