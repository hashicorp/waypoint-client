// Code generated by go-swagger; DO NOT EDIT.

package github_control_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint-client/gen/models"
)

// GithubControlServiceOAuthCallbackReader is a Reader for the GithubControlServiceOAuthCallback structure.
type GithubControlServiceOAuthCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GithubControlServiceOAuthCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGithubControlServiceOAuthCallbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGithubControlServiceOAuthCallbackDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGithubControlServiceOAuthCallbackOK creates a GithubControlServiceOAuthCallbackOK with default headers values
func NewGithubControlServiceOAuthCallbackOK() *GithubControlServiceOAuthCallbackOK {
	return &GithubControlServiceOAuthCallbackOK{}
}

/*GithubControlServiceOAuthCallbackOK handles this case with default header values.

A successful response.
*/
type GithubControlServiceOAuthCallbackOK struct {
	Payload *models.HashicorpCloudWaypointOAuthCallbackResponse
}

func (o *GithubControlServiceOAuthCallbackOK) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-04-21/github/callback.json][%d] githubControlServiceOAuthCallbackOK  %+v", 200, o.Payload)
}

func (o *GithubControlServiceOAuthCallbackOK) GetPayload() *models.HashicorpCloudWaypointOAuthCallbackResponse {
	return o.Payload
}

func (o *GithubControlServiceOAuthCallbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudWaypointOAuthCallbackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGithubControlServiceOAuthCallbackDefault creates a GithubControlServiceOAuthCallbackDefault with default headers values
func NewGithubControlServiceOAuthCallbackDefault(code int) *GithubControlServiceOAuthCallbackDefault {
	return &GithubControlServiceOAuthCallbackDefault{
		_statusCode: code,
	}
}

/*GithubControlServiceOAuthCallbackDefault handles this case with default header values.

An unexpected error response.
*/
type GithubControlServiceOAuthCallbackDefault struct {
	_statusCode int

	Payload *models.GrpcGatewayRuntimeError
}

// Code gets the status code for the github control service o auth callback default response
func (o *GithubControlServiceOAuthCallbackDefault) Code() int {
	return o._statusCode
}

func (o *GithubControlServiceOAuthCallbackDefault) Error() string {
	return fmt.Sprintf("[GET /waypoint/2022-04-21/github/callback.json][%d] GithubControlService_OAuthCallback default  %+v", o._statusCode, o.Payload)
}

func (o *GithubControlServiceOAuthCallbackDefault) GetPayload() *models.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GithubControlServiceOAuthCallbackDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
