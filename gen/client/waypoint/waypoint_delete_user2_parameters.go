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
)

// NewWaypointDeleteUser2Params creates a new WaypointDeleteUser2Params object
// with the default values initialized.
func NewWaypointDeleteUser2Params() *WaypointDeleteUser2Params {
	var ()
	return &WaypointDeleteUser2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointDeleteUser2ParamsWithTimeout creates a new WaypointDeleteUser2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointDeleteUser2ParamsWithTimeout(timeout time.Duration) *WaypointDeleteUser2Params {
	var ()
	return &WaypointDeleteUser2Params{

		timeout: timeout,
	}
}

// NewWaypointDeleteUser2ParamsWithContext creates a new WaypointDeleteUser2Params object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointDeleteUser2ParamsWithContext(ctx context.Context) *WaypointDeleteUser2Params {
	var ()
	return &WaypointDeleteUser2Params{

		Context: ctx,
	}
}

// NewWaypointDeleteUser2ParamsWithHTTPClient creates a new WaypointDeleteUser2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointDeleteUser2ParamsWithHTTPClient(client *http.Client) *WaypointDeleteUser2Params {
	var ()
	return &WaypointDeleteUser2Params{
		HTTPClient: client,
	}
}

/*WaypointDeleteUser2Params contains all the parameters to send to the API endpoint
for the waypoint delete user2 operation typically these are written to a http.Request
*/
type WaypointDeleteUser2Params struct {

	/*UserIDID*/
	UserIDID *string
	/*UserUsernameUsername*/
	UserUsernameUsername string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) WithTimeout(timeout time.Duration) *WaypointDeleteUser2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) WithContext(ctx context.Context) *WaypointDeleteUser2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) WithHTTPClient(client *http.Client) *WaypointDeleteUser2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserIDID adds the userIDID to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) WithUserIDID(userIDID *string) *WaypointDeleteUser2Params {
	o.SetUserIDID(userIDID)
	return o
}

// SetUserIDID adds the userIdId to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) SetUserIDID(userIDID *string) {
	o.UserIDID = userIDID
}

// WithUserUsernameUsername adds the userUsernameUsername to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) WithUserUsernameUsername(userUsernameUsername string) *WaypointDeleteUser2Params {
	o.SetUserUsernameUsername(userUsernameUsername)
	return o
}

// SetUserUsernameUsername adds the userUsernameUsername to the waypoint delete user2 params
func (o *WaypointDeleteUser2Params) SetUserUsernameUsername(userUsernameUsername string) {
	o.UserUsernameUsername = userUsernameUsername
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointDeleteUser2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserIDID != nil {

		// query param user.id.id
		var qrUserIDID string
		if o.UserIDID != nil {
			qrUserIDID = *o.UserIDID
		}
		qUserIDID := qrUserIDID
		if qUserIDID != "" {
			if err := r.SetQueryParam("user.id.id", qUserIDID); err != nil {
				return err
			}
		}

	}

	// path param user.username.username
	if err := r.SetPathParam("user.username.username", o.UserUsernameUsername); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
