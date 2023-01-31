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

// NewWaypointGetOIDCAuthURLParams creates a new WaypointGetOIDCAuthURLParams object
// with the default values initialized.
func NewWaypointGetOIDCAuthURLParams() *WaypointGetOIDCAuthURLParams {
	var ()
	return &WaypointGetOIDCAuthURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetOIDCAuthURLParamsWithTimeout creates a new WaypointGetOIDCAuthURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointGetOIDCAuthURLParamsWithTimeout(timeout time.Duration) *WaypointGetOIDCAuthURLParams {
	var ()
	return &WaypointGetOIDCAuthURLParams{

		timeout: timeout,
	}
}

// NewWaypointGetOIDCAuthURLParamsWithContext creates a new WaypointGetOIDCAuthURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointGetOIDCAuthURLParamsWithContext(ctx context.Context) *WaypointGetOIDCAuthURLParams {
	var ()
	return &WaypointGetOIDCAuthURLParams{

		Context: ctx,
	}
}

// NewWaypointGetOIDCAuthURLParamsWithHTTPClient creates a new WaypointGetOIDCAuthURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointGetOIDCAuthURLParamsWithHTTPClient(client *http.Client) *WaypointGetOIDCAuthURLParams {
	var ()
	return &WaypointGetOIDCAuthURLParams{
		HTTPClient: client,
	}
}

/*WaypointGetOIDCAuthURLParams contains all the parameters to send to the API endpoint
for the waypoint get o ID c auth URL operation typically these are written to a http.Request
*/
type WaypointGetOIDCAuthURLParams struct {

	/*AuthMethodName*/
	AuthMethodName string
	/*Body*/
	Body *models.HashicorpWaypointGetOIDCAuthURLRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) WithTimeout(timeout time.Duration) *WaypointGetOIDCAuthURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) WithContext(ctx context.Context) *WaypointGetOIDCAuthURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) WithHTTPClient(client *http.Client) *WaypointGetOIDCAuthURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthMethodName adds the authMethodName to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) WithAuthMethodName(authMethodName string) *WaypointGetOIDCAuthURLParams {
	o.SetAuthMethodName(authMethodName)
	return o
}

// SetAuthMethodName adds the authMethodName to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) SetAuthMethodName(authMethodName string) {
	o.AuthMethodName = authMethodName
}

// WithBody adds the body to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) WithBody(body *models.HashicorpWaypointGetOIDCAuthURLRequest) *WaypointGetOIDCAuthURLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint get o ID c auth URL params
func (o *WaypointGetOIDCAuthURLParams) SetBody(body *models.HashicorpWaypointGetOIDCAuthURLRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetOIDCAuthURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param auth_method.name
	if err := r.SetPathParam("auth_method.name", o.AuthMethodName); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
