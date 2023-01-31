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

// WaypointGetPushedArtifact2Reader is a Reader for the WaypointGetPushedArtifact2 structure.
type WaypointGetPushedArtifact2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WaypointGetPushedArtifact2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWaypointGetPushedArtifact2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWaypointGetPushedArtifact2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWaypointGetPushedArtifact2OK creates a WaypointGetPushedArtifact2OK with default headers values
func NewWaypointGetPushedArtifact2OK() *WaypointGetPushedArtifact2OK {
	return &WaypointGetPushedArtifact2OK{}
}

/*WaypointGetPushedArtifact2OK handles this case with default header values.

A successful response.
*/
type WaypointGetPushedArtifact2OK struct {
	Payload *models.HashicorpWaypointPushedArtifact
}

func (o *WaypointGetPushedArtifact2OK) Error() string {
	return fmt.Sprintf("[GET /project/{ref.sequence.application.project}/application/{ref.sequence.application.application}/artifact/{ref.sequence.number}][%d] waypointGetPushedArtifact2OK  %+v", 200, o.Payload)
}

func (o *WaypointGetPushedArtifact2OK) GetPayload() *models.HashicorpWaypointPushedArtifact {
	return o.Payload
}

func (o *WaypointGetPushedArtifact2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpWaypointPushedArtifact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWaypointGetPushedArtifact2Default creates a WaypointGetPushedArtifact2Default with default headers values
func NewWaypointGetPushedArtifact2Default(code int) *WaypointGetPushedArtifact2Default {
	return &WaypointGetPushedArtifact2Default{
		_statusCode: code,
	}
}

/*WaypointGetPushedArtifact2Default handles this case with default header values.

An unexpected error response.
*/
type WaypointGetPushedArtifact2Default struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the waypoint get pushed artifact2 default response
func (o *WaypointGetPushedArtifact2Default) Code() int {
	return o._statusCode
}

func (o *WaypointGetPushedArtifact2Default) Error() string {
	return fmt.Sprintf("[GET /project/{ref.sequence.application.project}/application/{ref.sequence.application.application}/artifact/{ref.sequence.number}][%d] Waypoint_GetPushedArtifact2 default  %+v", o._statusCode, o.Payload)
}

func (o *WaypointGetPushedArtifact2Default) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *WaypointGetPushedArtifact2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
