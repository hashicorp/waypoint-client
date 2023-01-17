// Code generated by go-swagger; DO NOT EDIT.

package waypoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/waypoint-client/gen/models"
)

// NewWaypointRunPipeline2Params creates a new WaypointRunPipeline2Params object
// with the default values initialized.
func NewWaypointRunPipeline2Params() *WaypointRunPipeline2Params {
	var ()
	return &WaypointRunPipeline2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointRunPipeline2ParamsWithTimeout creates a new WaypointRunPipeline2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointRunPipeline2ParamsWithTimeout(timeout time.Duration) *WaypointRunPipeline2Params {
	var ()
	return &WaypointRunPipeline2Params{

		timeout: timeout,
	}
}

// NewWaypointRunPipeline2ParamsWithContext creates a new WaypointRunPipeline2Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointRunPipeline2ParamsWithContext(ctx context.Context) *WaypointRunPipeline2Params {
	var ()
	return &WaypointRunPipeline2Params{

		Context: ctx,
	}
}

// NewWaypointRunPipeline2ParamsWithHTTPClient creates a new WaypointRunPipeline2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointRunPipeline2ParamsWithHTTPClient(client *http.Client) *WaypointRunPipeline2Params {
	var ()
	return &WaypointRunPipeline2Params{
		HTTPClient: client,
	}
}

/*WaypointRunPipeline2Params contains all the parameters to send to the API endpoint
for the waypoint run pipeline2 operation typically these are written to a http.Request
*/
type WaypointRunPipeline2Params struct {

	/*Body*/
	Body *models.HashicorpWaypointRunPipelineRequest
	/*PipelineID
	  Reference a single pipeline by ID.

	*/
	PipelineID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) WithTimeout(timeout time.Duration) *WaypointRunPipeline2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) WithContext(ctx context.Context) *WaypointRunPipeline2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) WithHTTPClient(client *http.Client) *WaypointRunPipeline2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) WithBody(body *models.HashicorpWaypointRunPipelineRequest) *WaypointRunPipeline2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) SetBody(body *models.HashicorpWaypointRunPipelineRequest) {
	o.Body = body
}

// WithPipelineID adds the pipelineID to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) WithPipelineID(pipelineID string) *WaypointRunPipeline2Params {
	o.SetPipelineID(pipelineID)
	return o
}

// SetPipelineID adds the pipelineId to the waypoint run pipeline2 params
func (o *WaypointRunPipeline2Params) SetPipelineID(pipelineID string) {
	o.PipelineID = pipelineID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointRunPipeline2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param pipeline.id
	if err := r.SetPathParam("pipeline.id", o.PipelineID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
