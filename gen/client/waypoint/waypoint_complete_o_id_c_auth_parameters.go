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

// NewWaypointCompleteOIDCAuthParams creates a new WaypointCompleteOIDCAuthParams object
// with the default values initialized.
func NewWaypointCompleteOIDCAuthParams() *WaypointCompleteOIDCAuthParams {
	var ()
	return &WaypointCompleteOIDCAuthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointCompleteOIDCAuthParamsWithTimeout creates a new WaypointCompleteOIDCAuthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointCompleteOIDCAuthParamsWithTimeout(timeout time.Duration) *WaypointCompleteOIDCAuthParams {
	var ()
	return &WaypointCompleteOIDCAuthParams{

		timeout: timeout,
	}
}

// NewWaypointCompleteOIDCAuthParamsWithContext creates a new WaypointCompleteOIDCAuthParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointCompleteOIDCAuthParamsWithContext(ctx context.Context) *WaypointCompleteOIDCAuthParams {
	var ()
	return &WaypointCompleteOIDCAuthParams{

		Context: ctx,
	}
}

// NewWaypointCompleteOIDCAuthParamsWithHTTPClient creates a new WaypointCompleteOIDCAuthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointCompleteOIDCAuthParamsWithHTTPClient(client *http.Client) *WaypointCompleteOIDCAuthParams {
	var ()
	return &WaypointCompleteOIDCAuthParams{
		HTTPClient: client,
	}
}

/*WaypointCompleteOIDCAuthParams contains all the parameters to send to the API endpoint
for the waypoint complete o ID c auth operation typically these are written to a http.Request
*/
type WaypointCompleteOIDCAuthParams struct {

	/*AuthMethodName*/
	AuthMethodName string
	/*Body*/
	Body *models.HashicorpWaypointCompleteOIDCAuthRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) WithTimeout(timeout time.Duration) *WaypointCompleteOIDCAuthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) WithContext(ctx context.Context) *WaypointCompleteOIDCAuthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) WithHTTPClient(client *http.Client) *WaypointCompleteOIDCAuthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthMethodName adds the authMethodName to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) WithAuthMethodName(authMethodName string) *WaypointCompleteOIDCAuthParams {
	o.SetAuthMethodName(authMethodName)
	return o
}

// SetAuthMethodName adds the authMethodName to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) SetAuthMethodName(authMethodName string) {
	o.AuthMethodName = authMethodName
}

// WithBody adds the body to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) WithBody(body *models.HashicorpWaypointCompleteOIDCAuthRequest) *WaypointCompleteOIDCAuthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint complete o ID c auth params
func (o *WaypointCompleteOIDCAuthParams) SetBody(body *models.HashicorpWaypointCompleteOIDCAuthRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointCompleteOIDCAuthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
