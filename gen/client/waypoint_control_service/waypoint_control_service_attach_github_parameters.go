// Code generated by go-swagger; DO NOT EDIT.

package waypoint_control_service

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

// NewWaypointControlServiceAttachGithubParams creates a new WaypointControlServiceAttachGithubParams object
// with the default values initialized.
func NewWaypointControlServiceAttachGithubParams() *WaypointControlServiceAttachGithubParams {
	var ()
	return &WaypointControlServiceAttachGithubParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointControlServiceAttachGithubParamsWithTimeout creates a new WaypointControlServiceAttachGithubParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWaypointControlServiceAttachGithubParamsWithTimeout(timeout time.Duration) *WaypointControlServiceAttachGithubParams {
	var ()
	return &WaypointControlServiceAttachGithubParams{

		timeout: timeout,
	}
}

// NewWaypointControlServiceAttachGithubParamsWithContext creates a new WaypointControlServiceAttachGithubParams object
// with the default values initialized, and the ability to set a context for a request
func NewWaypointControlServiceAttachGithubParamsWithContext(ctx context.Context) *WaypointControlServiceAttachGithubParams {
	var ()
	return &WaypointControlServiceAttachGithubParams{

		Context: ctx,
	}
}

// NewWaypointControlServiceAttachGithubParamsWithHTTPClient creates a new WaypointControlServiceAttachGithubParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWaypointControlServiceAttachGithubParamsWithHTTPClient(client *http.Client) *WaypointControlServiceAttachGithubParams {
	var ()
	return &WaypointControlServiceAttachGithubParams{
		HTTPClient: client,
	}
}

/*WaypointControlServiceAttachGithubParams contains all the parameters to send to the API endpoint
for the waypoint control service attach github operation typically these are written to a http.Request
*/
type WaypointControlServiceAttachGithubParams struct {

	/*GithubAuthToken*/
	GithubAuthToken *string
	/*NamespaceID*/
	NamespaceID string
	/*Project*/
	Project *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) WithTimeout(timeout time.Duration) *WaypointControlServiceAttachGithubParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) WithContext(ctx context.Context) *WaypointControlServiceAttachGithubParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) WithHTTPClient(client *http.Client) *WaypointControlServiceAttachGithubParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGithubAuthToken adds the githubAuthToken to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) WithGithubAuthToken(githubAuthToken *string) *WaypointControlServiceAttachGithubParams {
	o.SetGithubAuthToken(githubAuthToken)
	return o
}

// SetGithubAuthToken adds the githubAuthToken to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) SetGithubAuthToken(githubAuthToken *string) {
	o.GithubAuthToken = githubAuthToken
}

// WithNamespaceID adds the namespaceID to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) WithNamespaceID(namespaceID string) *WaypointControlServiceAttachGithubParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithProject adds the project to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) WithProject(project *string) *WaypointControlServiceAttachGithubParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the waypoint control service attach github params
func (o *WaypointControlServiceAttachGithubParams) SetProject(project *string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointControlServiceAttachGithubParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GithubAuthToken != nil {

		// query param github_auth_token
		var qrGithubAuthToken string
		if o.GithubAuthToken != nil {
			qrGithubAuthToken = *o.GithubAuthToken
		}
		qGithubAuthToken := qrGithubAuthToken
		if qGithubAuthToken != "" {
			if err := r.SetQueryParam("github_auth_token", qGithubAuthToken); err != nil {
				return err
			}
		}

	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if o.Project != nil {

		// query param project
		var qrProject string
		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {
			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}